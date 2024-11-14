package data

import "github.com/proyectos01-a/RestaurantMenu/src/models"

type ChatHistoryRepository interface {
	SaveChatHistory(senderWspNumber string, message string, botResponse string) error
	GetChatHistory(senderWspNumber string) ([]models.ChatHistory, error)
}
