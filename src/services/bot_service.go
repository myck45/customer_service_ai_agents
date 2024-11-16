package services

import (
	"context"

	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/sashabaranov/go-openai"
)

type BotService interface {
	GenerateEmbedding(data string) ([]float32, error)
	GenerateBotResponse(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error)
	BotResponse(chat request.TwilioWebhook) error
	SystemPrompt(aditionalData string) (string, error)
	PrepareChatMessages(chat request.TwilioWebhook, semanticContext []response.MenuSearchResponse, restaurantID uint) ([]openai.ChatCompletionMessage, error)
	TwilioResponse(userWspNumber string, botWspNumber string, botResponse string) error
}
