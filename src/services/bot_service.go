package services

import "github.com/proyectos01-a/RestaurantMenu/src/models"

type BotService interface {
	GenerateEmbedding(data string) ([]float32, error)
	ChatCompletion(data string) (string, error)
	SystemPrompt(aditionalData string) (string, error)
	GetChatHistory(senderWspNumber string) ([]models.ChatHistory, error)
	SaveChatHistory(senderWspNumber string, message string, botResponse string) error
}
