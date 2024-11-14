package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/proyectos01-a/RestaurantMenu/src/data"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type BotServiceImpl struct {
	openAI          *openai.Client
	chatHistoryRepo data.ChatHistoryRepository
}

// GetChatHistory implements BotService.
func (b *BotServiceImpl) GetChatHistory(senderWspNumber string) ([]models.ChatHistory, error) {

	chatHistory, err := b.chatHistoryRepo.GetChatHistory(senderWspNumber)
	if err != nil {
		logrus.WithError(err).Error("failed to get chat history")
		return nil, err
	}

	return chatHistory, nil
}

// SaveChatHistory implements BotService.
func (b *BotServiceImpl) SaveChatHistory(senderWspNumber string, message string, botResponse string) error {

	err := b.chatHistoryRepo.SaveChatHistory(senderWspNumber, message, botResponse)
	if err != nil {
		logrus.WithError(err).Error("failed to save chat history")
		return err
	}

	return nil
}

// ChatCompletion implements BotService.
func (b *BotServiceImpl) ChatCompletion(data string) (string, error) {
	panic("implement me")

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

func NewBotServiceImpl(openAI *openai.Client, chatHistoryRepo data.ChatHistoryRepository) BotService {
	return &BotServiceImpl{
		openAI:          openAI,
		chatHistoryRepo: chatHistoryRepo,
	}
}
