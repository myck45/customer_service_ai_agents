package data

import "github.com/proyectos01-a/shared/models"

type ChatHistoryRepository interface {
	SaveChat(chatHistory *models.ChatHistory) error
	GetChatHistoryBySenderWspNumberAndRestaurantID(senderWspNumber string, restaurantID uint) ([]models.ChatHistory, error)
	GetChatHistory(senderWspNumber string, botWspNumber string, restaurantID uint) ([]models.ChatHistory, error)
}
