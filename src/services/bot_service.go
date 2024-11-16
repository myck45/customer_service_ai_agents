package services

import (
	"context"

	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/sashabaranov/go-openai"
)

type BotService interface {

	// Bot CRUD methods
	GetBotByID(botID uint) (*response.BotResponse, error)
	GetBotByRestaurantID(restaurantID uint) ([]response.BotResponse, error)
	GetBotByWspNumber(wspNumber string) (*response.BotResponse, error)
	GetAllBots() ([]response.BotResponse, error)
	CreateBot(bot request.CreateBotReq) (*response.BotResponse, error)
	UpdateBot(botID uint, bot request.UpdateBotReq) (*response.BotResponse, error)
	DeleteBotByID(botID uint) error

	// Bot response methods
	GenerateEmbedding(data string) ([]float32, error)
	GenerateBotResponse(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error)
	BotResponse(chat request.TwilioWebhook) error
	SystemPrompt(botConfig request.BotConfig) (string, error)
	PrepareChatMessages(chat request.TwilioWebhook, semanticContext []response.MenuSearchResponse, botInfo request.BotInfo) ([]openai.ChatCompletionMessage, error)
	TwilioResponse(userWspNumber string, botWspNumber string, botResponse string) error
}
