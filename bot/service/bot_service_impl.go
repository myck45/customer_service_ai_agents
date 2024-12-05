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
			},
		},
	)
	if err != nil {
		logrus.WithError(err).Error("failed to create chat completion")
		return "", err
	}

	if res.Choices[0].FinishReason == openai.FinishReasonToolCalls {
		toolCall := res.Choices[0].Message.ToolCalls[0]
		args := toolCall.Function.Arguments
		orderCode, err := b.botToolHandler.HandleGetUserOrder(args, chatInfo)
		if err != nil {
			logrus.WithError(err).Error("failed to handle user order")
			return "", err
		}

		botResponse := fmt.Sprintf("Tu pedido ha sido registrado con éxito 🎉🍽️. Tu código de pedido es: %s. ¡Gracias por tu preferencia! 🙏✨", orderCode)

		return botResponse, nil
	}

	botResponse := res.Choices[0].Message.Content

	return botResponse, nil
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

	return messages, nil
}

// SystemPrompt implements BotService.
func (b *BotServiceImpl) SystemPrompt(botConfig req.BotConfig) (string, error) {

	additionalData := botConfig.SemanticContext
	botName := botConfig.BotName
	botIdentity := botConfig.BotIdentity

	systemPrompt := fmt.Sprintf(`
**Identidad**
- **Nombre** tu nombre es %s
- **Identidad** %s 
		
		
		Proporcionas información detallada sobre el menú, platos, y datos clave del restaurante usando un sistema de búsqueda semántica que enriquece las respuestas con contexto relevante.

**Capacidades y Comportamiento:**
- Respondes de forma clara y amigable, ajustándote a la consulta del usuario.
- Proporcionas detalles de platillos (ingredientes, preparación, alérgenos) y sugieres opciones similares si la similitud es alta.
- Das información sobre el restaurante (horarios, ubicación, estilo de cocina).
- Eres alegre, educado y respetuoso en todo momento, puedes usar emojis para expresarte mejor si es necesario.
- Tu personalidad es amigable y servicial, siempre buscas ayudar a los clientes.
- Eres persuasivo y promueves la calidad de los platillos y la experiencia en el restaurante.

**Uso de Búsqueda Semántica:**
- **Contexto Actual:** %s
- **Fecha:** %s
- Seleccionas platillos según similitud sin mencionar "contexto" o "grado de similitud". Si la consulta no requiere contexto, respondes de forma directa.
- Utiliza el contexto para enriquecer tus respuestas, pero no lo menciones explícitamente.

**Limitaciones y Directrices:**
1. No inventes datos ni reveles información confidencial.
2. Redirige temas fuera del restaurante hacia temas relevantes.
3. Tus respuestas son enviadas por WhatsApp, por lo que debes adaptar el formato de tus respuestas a mensajes que puedan ser presentados en esa plataforma.

**Objetivo:** 
Ofrecer una experiencia informativa y accesible para que los clientes conozcan más sobre el restaurante y su menú, promoviendo satisfacción e interés.

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
