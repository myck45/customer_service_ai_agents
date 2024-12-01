package service

import (
	"context"

	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/models"
	"github.com/sashabaranov/go-openai"
)

type BotService interface {
	GenerateBotResponse(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error)
	BotResponse(chat *req.TwilioWebhook) error
	SystemPrompt(botConfig req.BotConfig) (string, error)
	PrepareChatMessages(chatHistory []models.ChatHistory, semanticContext []dto.MenuSearchResponse, botInfo req.BotInfo) ([]openai.ChatCompletionMessage, error)
}
