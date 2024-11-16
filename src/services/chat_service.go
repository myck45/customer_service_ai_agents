package services

import (
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
)

type ChatService interface {
	SaveChat(chat request.TwilioWebhook, botResponse string) error
	GetChatHistoryBySenderWspNumber(senderWspNumber string) ([]models.ChatHistory, error)
}
