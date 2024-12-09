package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/shared/models"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type BotServiceImpl struct {
	openAI          *openai.Client
	twilio          utils.TwilioUtils
	botUtils        utils.BotUtils
	chatHistoryRepo data.ChatHistoryRepository
	botRepo         data.BotRepository
	menuRepo        data.MenuRepository
	botTools        utils.BotTools
	botToolHandler  handlers.BotToolsHandler
}

const (
	botErrResp = "Lo siento 😔💔, por el momento no puedo responder a tu consulta 🤖❌. Por favor, intenta de nuevo más tarde 🙏✨."
)

// BotResponse implements BotService.
func (b *BotServiceImpl) BotResponse(chat *req.TwilioWebhook) error {

	var similarityThreshold float32 = 0.5 // minimum similarity threshold for the semantic search
	var matchCount int = 5                // is the number of menu items that the query will return
	var botWspNumber string = chat.To     // is the WhatsApp number of the bot
	var userWspNumber string = chat.From  // is the WhatsApp number of the user
	var userMessage string = chat.Body    // is the message sent by the user

	if !strings.HasPrefix(botWspNumber, "whatsapp:") {
		botWspNumber = fmt.Sprintf("whatsapp:%s", botWspNumber)
	}

	// Get the bot by the WhatsApp number
	bot, err := b.botRepo.GetBotByWspNumber(botWspNumber)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to get bot: %v", err)
	}

	botInfo := req.BotInfo{
		BotName:      bot.Name,
		BotIdentity:  bot.Identity,
		RestaurantID: bot.RestaurantID,
	}

	// Generate the embedding for the user message
	userMsgEmbedding, err := b.botUtils.GenerateEmbedding(userMessage)
	if err != nil {
		logrus.WithError(err).Error("failed to generate embedding")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to generate embedding: %v", err)
	}

	// Search the menu using the semantic context
	semanticContext, err := b.menuRepo.SemanticSearchMenu(userMsgEmbedding, similarityThreshold, matchCount, botInfo.RestaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to search menu")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to search menu: %v", err)
	}

	// Get the chat history
	chatHistory, err := b.chatHistoryRepo.GetChatHistory(userWspNumber, botWspNumber, botInfo.RestaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to get chat history")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to get chat history: %v", err)
	}

	// Prepare the chat messages
	messages, err := b.PrepareChatMessages(chatHistory, semanticContext, botInfo)
	if err != nil {
		logrus.WithError(err).Error("failed to prepare chat messages")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to prepare chat messages: %v", err)
	}

	// Append the current user message
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: userMessage,
	})

	chatInfo := dto.ChatInfoRequest{
		BotWspNumber:    botWspNumber,
		SenderWspNumber: userWspNumber,
		RestaurantID:    botInfo.RestaurantID,
	}

	// Generate the bot response
	botResponse, err := b.GenerateBotResponse(context.Background(), messages, chatInfo)
	if err != nil {
		logrus.WithError(err).Error("failed to generate bot response")
		if twErr := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botErrResp); twErr != nil {
			logrus.WithError(twErr).Error("failed to send response")
		}
		return fmt.Errorf("failed to generate bot response: %v", err)
	}

	// Send the bot response through Twilio
	if err := b.twilio.SendWspMessage(userWspNumber, botWspNumber, botResponse); err != nil {
		logrus.WithError(err).Error("failed to send response")
		return fmt.Errorf("failed to send response: %v", err)
	}

	// Save the chat history
	err = b.chatHistoryRepo.SaveChat(&models.ChatHistory{
		SenderWspNumber: userWspNumber,
		BotWspNumber:    botWspNumber,
		Message:         userMessage,
		BotResponse:     botResponse,
		RestaurantID:    botInfo.RestaurantID,
	})
	if err != nil {
		logrus.WithError(err).Error("failed to save chat")
		return fmt.Errorf("failed to save chat: %v", err)
	}

	return nil
}

// GenerateBotResponse implements BotService.
func (b *BotServiceImpl) GenerateBotResponse(ctx context.Context, messages []openai.ChatCompletionMessage, chatInfo dto.ChatInfoRequest) (string, error) {

	// Create the chat completion request
	res, err := b.openAI.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: messages,
			Tools: []openai.Tool{
				{
					Type:     openai.ToolTypeFunction,
					Function: b.botTools.GetUserOrder(),
				},
				{
					Type:     openai.ToolTypeFunction,
					Function: b.botTools.DeleteUserOrder(),
				},
			},
		},
	)
	if err != nil {
		logrus.WithError(err).Error("failed to create chat completion")
		return "", err
	}

	if res.Choices[0].FinishReason == openai.FinishReasonToolCalls {
		toolCall := res.Choices[0].Message.ToolCalls[0]
		botResponse, err := b.HandleBotToolCall(toolCall, chatInfo)
		if err != nil {
			logrus.WithError(err).Error("failed to handle bot tool call")
			return "", err
		}
		return botResponse, nil

	}

	botResponse := res.Choices[0].Message.Content

	return botResponse, nil
}

// HandleBotToolCall implements BotService.
func (b *BotServiceImpl) HandleBotToolCall(toolCall openai.ToolCall, chatInfo dto.ChatInfoRequest) (string, error) {

	functionName := toolCall.Function.Name
	args := toolCall.Function.Arguments

	if functionName == "get_user_order" {
		order, err := b.botToolHandler.HandleGetUserOrder(args, chatInfo)
		if err != nil {
			logrus.WithError(err).Error("failed to handle user order")
			return "", err
		}

		var details string
		for _, item := range order.OrderMenuItems {
			details += fmt.Sprintf("- %s (x%d) $%d\n", item.ItemName, item.Quantity, item.Subtotal)
		}

		botResponse := fmt.Sprintf(
			"🍔🍟 *Pedido Realizado* 🍟🍔\n\n"+
				"*Detalles del Pedido:*\n\n"+
				"%s"+
				"\n"+
				"*Código único de tu pedido:*\n\n"+
				"- %s\n\n"+
				"_este código es importante para rastrear tu pedido o para cancelarlo._\n\n"+
				"*Dirección de Entrega*: %s\n"+
				"*Método de Pago*: %s\n"+
				"*Total*: $%d\n\n"+
				"🛵 ¡Tu pedido está en camino! 🛵\n"+
				"🍽️ ¡Gracias por tu compra! 🍽️",
			details, order.OrderCode, order.DeliveryAddress, order.PaymentMethod, order.TotalPrice,
		)

		return botResponse, nil
	}

	if functionName == "delete_user_order" {
		orderCode, err := b.botToolHandler.HandleDeleteUserOrder(args, chatInfo)
		if err != nil {
			logrus.WithError(err).Error("failed to handle delete user order")
			return "", err
		}

		botResponse := fmt.Sprintf(
			"🚫 ¡Tu pedido ha sido cancelado! 🚫\n\n"+
				"*El pedido con Código: %s a sido cancelado*\n\n"+
				"🍟 ¡Gracias por tu preferencia! 🍟",
			orderCode,
		)

		return botResponse, nil
	}

	defaultResponse := fmt.Sprintf(
		"Hubo un error al procesar tu solicitud, por favor intenta de nuevo más tarde. 😔💔\n\n" +
			"Si necesitas ayuda, por favor comunícate con nuestro equipo de soporte. 📞📧",
	)

	return defaultResponse, nil
}

// PrepareChatMessages implements BotService.
func (b *BotServiceImpl) PrepareChatMessages(chatHistory []models.ChatHistory, semanticContext []dto.MenuSearchResponse, botInfo req.BotInfo) ([]openai.ChatCompletionMessage, error) {

	contextStr, err := json.Marshal(semanticContext)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal semantic context")
		return nil, err
	}

	botConfig := req.BotConfig{
		BotName:         botInfo.BotName,
		BotIdentity:     botInfo.BotIdentity,
		SemanticContext: string(contextStr),
	}

	systemPrompt, err := b.SystemPrompt(botConfig)
	if err != nil {
		logrus.WithError(err).Error("failed to generate system prompt")
		return nil, err
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	// Invert the chat history
	for i, j := 0, len(chatHistory)-1; i < j; i, j = i+1, j-1 {
		chatHistory[i], chatHistory[j] = chatHistory[j], chatHistory[i]
	}

	for _, chat := range chatHistory {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: chat.Message,
		})
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: chat.BotResponse,
		})
	}

	logrus.Infof("messages: %+v", messages)

	return messages, nil
}

// SystemPrompt implements BotService.
func (b *BotServiceImpl) SystemPrompt(botConfig req.BotConfig) (string, error) {

	additionalData := botConfig.SemanticContext
	botName := botConfig.BotName
	botIdentity := botConfig.BotIdentity

	systemPrompt := fmt.Sprintf(`
# **Identidad**
- **Nombre**: tu nombre es %s
- **Identidad**: %s 

Proporcionas información detallada sobre el menú, platos, y datos clave del restaurante usando un sistema de búsqueda semántica que enriquece las respuestas con contexto relevante.

**Capacidades y Comportamiento:**
- Respondes de forma clara y amigable, ajustándote a la consulta del usuario.
- Proporcionas detalles de platillos (ingredientes, preparación, alérgenos) y sugieres opciones similares si la similitud es alta.
- Promueves el restaurante, hablas bien de este, tratando siempre de atraer a los clientes.
- Eres alegre, educado y respetuoso en todo momento, puedes usar emojis para expresarte mejor si es necesario.
- Tu personalidad es amigable y servicial, siempre buscas ayudar a los clientes.
- Eres persuasivo y promueves la calidad de los platillos y la experiencia en el restaurante.
- Tienes la capacidad de registrar pedidos o  de cancelar los pedidos, pero no de actualizarlos, debes indicar esto al usuario cuando sea necesario, para que el usuario lo piense bien antes de registrar un pedido.
- Solo puedes recordar las últimas 5 interacciones con el usuario, es decir 5 mensajes enviados por el usuario y 5 mensajes enviados por ti, para que lo tengas en cuenta. Mencionalo solo si es necesario.

**Contexto actual:**
- **Menú Disponible:** %s
- **Fecha Actual:** %s
- Seleccionas platillos según similitud sin mencionar "contexto" o "grado de similitud". Si la consulta no requiere contexto, respondes de forma directa.
- Utiliza el contexto para enriquecer tus respuestas, pero no lo menciones explícitamente.
- El contexto son los platillos disponibles en el menú. **Solo puedes ofrecer al cliente los platillos disponibles en el menú.**

**Limitaciones y Directrices:**
1. No inventes datos ni reveles información confidencial.
2. Redirige temas fuera del restaurante hacia temas relevantes.
3. Tus respuestas son enviadas por WhatsApp, por lo que debes adaptar el formato de tus respuestas a mensajes que puedan ser presentados en esa plataforma.
4. Tienes estrictamente prohibido ofrecer información que no esté relacionada con el restaurante o el menú.
5. **Solo puedes ofrecer platillos que están en el menú actual. No inventes platillos.**

**Ejemplos de Respuestas:**
- Si un cliente pregunta por un platillo específico, responde con los detalles de ese platillo si está en el menú.
- Si un cliente pregunta por recomendaciones, sugiere platillos del menú actual.
- Si un cliente pregunta por el menú disponible, habla solo de los platillos presentes en el contexto actual.

**Debes seguir estas directrices para garantizar una experiencia de usuario óptima, de lo contrario serás despedido.**
    `, botName, botIdentity, additionalData, time.Now().Format("2006-01-02"))

	return systemPrompt, nil
}

func NewBotServiceImpl(openAI *openai.Client, twilio utils.TwilioUtils, botUtils utils.BotUtils, chatHistoryRepo data.ChatHistoryRepository, botRepo data.BotRepository, menuRepo data.MenuRepository, botTools utils.BotTools, botToolHandler handlers.BotToolsHandler) BotService {
	return &BotServiceImpl{
		openAI:          openAI,
		twilio:          twilio,
		botUtils:        botUtils,
		chatHistoryRepo: chatHistoryRepo,
		botRepo:         botRepo,
		menuRepo:        menuRepo,
		botTools:        botTools,
		botToolHandler:  botToolHandler,
	}
}
