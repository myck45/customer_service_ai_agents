package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/proyectos01-a/RestaurantMenu/src/data"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type BotServiceImpl struct {
	openAI          *openai.Client
	twilio          *twilio.RestClient
	chatHistoryRepo data.ChatHistoryRepository
	botRepo         data.BotRepository
	menuRepo        data.MenuRepository
}

// CreateBot implements BotService.
func (b *BotServiceImpl) CreateBot(bot request.CreateBotReq) (*response.BotResponse, error) {

	// Create the bot
	newBot := &models.Bot{
		Name:         bot.Name,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	// Save the bot
	if err := b.botRepo.SaveBot(newBot); err != nil {
		logrus.WithError(err).Error("failed to save bot")
		return nil, fmt.Errorf("failed to save bot: %v", err)
	}

	// Create the response
	botResponse := &response.BotResponse{
		ID:           newBot.ID,
		Name:         newBot.Name,
		WspNumber:    newBot.WspNumber,
		RestaurantID: newBot.RestaurantID,
	}

	return botResponse, nil
}

// DeleteBotByID implements BotService.
func (b *BotServiceImpl) DeleteBotByID(botID uint) error {

	// Delete the bot
	if err := b.botRepo.DeleteBotByID(botID); err != nil {
		logrus.WithError(err).Error("failed to delete bot")
		return fmt.Errorf("failed to delete bot: %v", err)
	}

	return nil
}

// GetAllBots implements BotService.
func (b *BotServiceImpl) GetAllBots() ([]response.BotResponse, error) {

	// Get all bots
	bots, err := b.botRepo.GetAllBots()
	if err != nil {
		logrus.WithError(err).Error("failed to get all bots")
		return nil, fmt.Errorf("failed to get all bots: %v", err)
	}

	// Create the response
	var botResponses []response.BotResponse
	for _, bot := range bots {
		botResponses = append(botResponses, response.BotResponse{
			ID:           bot.ID,
			Name:         bot.Name,
			WspNumber:    bot.WspNumber,
			RestaurantID: bot.RestaurantID,
		})
	}

	return botResponses, nil
}

// GetBotByID implements BotService.
func (b *BotServiceImpl) GetBotByID(botID uint) (*response.BotResponse, error) {

	// Get the bot
	bot, err := b.botRepo.GetBotByID(botID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	botResponse := &response.BotResponse{
		ID:           bot.ID,
		Name:         bot.Name,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	return botResponse, nil
}

// GetBotByRestaurantID implements BotService.
func (b *BotServiceImpl) GetBotByRestaurantID(restaurantID uint) ([]response.BotResponse, error) {

	// Get the bots
	bots, err := b.botRepo.GetBotByRestaurantID(restaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	var botResponses []response.BotResponse
	for _, bot := range bots {
		botResponses = append(botResponses, response.BotResponse{
			ID:           bot.ID,
			Name:         bot.Name,
			WspNumber:    bot.WspNumber,
			RestaurantID: bot.RestaurantID,
		})
	}

	return botResponses, nil
}

// GetBotByWspNumber implements BotService.
func (b *BotServiceImpl) GetBotByWspNumber(wspNumber string) (*response.BotResponse, error) {

	// Get the bot
	bot, err := b.botRepo.GetBotByWspNumber(wspNumber)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	botResponse := &response.BotResponse{
		ID:           bot.ID,
		Name:         bot.Name,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	return botResponse, nil
}

// UpdateBot implements BotService.
func (b *BotServiceImpl) UpdateBot(botID uint, bot request.UpdateBotReq) (*response.BotResponse, error) {

	// Get the bot
	botToUpdate, err := b.botRepo.GetBotByID(botID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Update the bot
	botToUpdate.Name = bot.Name
	botToUpdate.WspNumber = bot.WspNumber

	// Save the bot
	if err := b.botRepo.SaveBot(botToUpdate); err != nil {
		logrus.WithError(err).Error("failed to save bot")
		return nil, fmt.Errorf("failed to save bot: %v", err)
	}

	// Create the response
	botResponse := &response.BotResponse{
		ID:           botToUpdate.ID,
		Name:         botToUpdate.Name,
		WspNumber:    botToUpdate.WspNumber,
		RestaurantID: botToUpdate.RestaurantID,
	}

	return botResponse, nil
}

// BotResponse implements BotService.
func (b *BotServiceImpl) BotResponse(chat request.TwilioWebhook) error {

	var similarityThreshold float32 = 0.5 // minimum similarity threshold for the semantic search
	var matchCount int = 5                // is the number of menu items that the query will return
	var botWspNumber string = chat.To     // is the WhatsApp number of the bot
	var userMessage string = chat.Body    // is the message sent by the user
	var userWspNumber string = chat.From  // is the WhatsApp number of the user

	// Get the bot by the WhatsApp number
	bot, err := b.botRepo.GetBotByWspNumber(botWspNumber)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return fmt.Errorf("failed to get bot: %v", err)
	}

	restaurantID := bot.RestaurantID

	// Generate the embedding for the user message
	userMsgEmbedding, err := b.GenerateEmbedding(userMessage)
	if err != nil {
		logrus.WithError(err).Error("failed to generate embedding")
		return fmt.Errorf("failed to generate embedding: %v", err)
	}

	// Search the menu using the semantic context
	semanticContext, err := b.menuRepo.SemanticSearchMenu(userMsgEmbedding, similarityThreshold, matchCount, restaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to search menu")
		return fmt.Errorf("failed to search menu: %v", err)
	}

	// Prepare the chat messages
	messages, err := b.PrepareChatMessages(chat, semanticContext, restaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to prepare chat messages")
		return fmt.Errorf("failed to prepare chat messages: %v", err)
	}

	// Generate the bot response
	botResponse, err := b.GenerateBotResponse(context.Background(), messages)
	if err != nil {
		logrus.WithError(err).Error("failed to generate bot response")
		return fmt.Errorf("failed to generate bot response: %v", err)
	}

	// Send the bot response through Twilio
	if err := b.TwilioResponse(userWspNumber, botWspNumber, botResponse); err != nil {
		logrus.WithError(err).Error("failed to send response")
		return fmt.Errorf("failed to send response: %v", err)
	}

	// Save the chat history
	err = b.chatHistoryRepo.SaveChat(&models.ChatHistory{
		SenderWspNumber: userWspNumber,
		BotWspNumber:    botWspNumber,
		Message:         userMessage,
		BotResponse:     botResponse,
		RestaurantID:    restaurantID,
	})
	if err != nil {
		logrus.WithError(err).Error("failed to save chat")
		return fmt.Errorf("failed to save chat: %v", err)
	}

	return nil

}

// GenerateBotResponse implements BotService.
func (b *BotServiceImpl) GenerateBotResponse(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	// Create the chat completion request
	res, err := b.openAI.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: messages,
		},
	)
	if err != nil {
		logrus.WithError(err).Error("failed to create chat completion")
		return "", err
	}

	// Get the bot response
	botResponse := res.Choices[0].Message.Content

	return botResponse, nil
}

// TwilioResponse implements BotService.
func (b *BotServiceImpl) TwilioResponse(userWspNumber string, botWspNumber string, botResponse string) error {

	userWspNumber = fmt.Sprintf("whatsapp:%s", userWspNumber)
	botWspNumber = fmt.Sprintf("whatsapp:%s", botWspNumber)

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(userWspNumber)
	params.SetFrom(botWspNumber)
	params.SetBody(botResponse)

	resp, err := b.twilio.Api.CreateMessage(params)
	if err != nil {
		logrus.WithError(err).Error("failed to send message")
		return fmt.Errorf("failed to send message: %v", err)
	}

	logrus.WithField("response", resp).Info("message sent")

	return nil
}

// PrepareChatMessages implements BotService.
func (b *BotServiceImpl) PrepareChatMessages(chat request.TwilioWebhook, semanticContext []response.MenuSearchResponse, restaurantID uint) ([]openai.ChatCompletionMessage, error) {

	var senderWspNumber string = chat.From
	var botWspNumber string = chat.To

	contextStr, err := json.Marshal(semanticContext)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal semantic context")
		return nil, err
	}

	systemPrompt, err := b.SystemPrompt(string(contextStr))
	if err != nil {
		logrus.WithError(err).Error("failed to generate system prompt")
		return nil, err
	}

	chatHistory, err := b.chatHistoryRepo.GetChatHistory(senderWspNumber, botWspNumber, restaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to get chat history")
		return nil, err
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: systemPrompt,
		},
	}

	for _, chat := range chatHistory {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: chat.Message,
		})
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: chat.BotResponse,
		})
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: chat.Body,
	})

	return messages, nil
}

// GenerateEmbedding implements BotService.
func (b *BotServiceImpl) GenerateEmbedding(data string) ([]float32, error) {

	if data == "" {
		return nil, errors.New("input data cannot be empty")
	}

	ctx := context.Background()

	targetReq := openai.EmbeddingRequest{
		Input: []string{data},
		Model: openai.LargeEmbedding3,
	}

	res, err := b.openAI.CreateEmbeddings(ctx, targetReq)
	if err != nil {
		logrus.WithError(err).Error("failed to create embeddings")
		return nil, err
	}

	embedding := res.Data[0].Embedding

	return embedding, nil
}

// SystemPrompt implements BotService.
func (b *BotServiceImpl) SystemPrompt(additionalData string) (string, error) {
	if additionalData == "" {
		return "", errors.New("additional data cannot be empty")
	}

	systemPrompt := fmt.Sprintf(`
		Eres un asistente de IA diseñado para mejorar la experiencia de los clientes en un restaurante. Proporcionas información detallada sobre el menú, platos, y datos clave del restaurante usando un sistema de búsqueda semántica que enriquece las respuestas con contexto relevante.

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

	`, additionalData, time.Now().Format("02-01-2006"))

	return systemPrompt, nil
}

func NewBotServiceImpl(openAI *openai.Client, twilio *twilio.RestClient, chatHistoryRepo data.ChatHistoryRepository, botRepo data.BotRepository, menuRepo data.MenuRepository) BotService {
	return &BotServiceImpl{
		openAI:          openAI,
		twilio:          twilio,
		chatHistoryRepo: chatHistoryRepo,
		botRepo:         botRepo,
		menuRepo:        menuRepo,
	}
}
