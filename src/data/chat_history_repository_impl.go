package data

import (
	"errors"

	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChatHistoryRepositoryImpl struct {
	db *gorm.DB
}

// GetChatHistory implements ChatHistoryRepository.
func (c *ChatHistoryRepositoryImpl) GetChatHistory(senderWspNumber string) ([]models.ChatHistory, error) {

	var chatHistory []models.ChatHistory

	result := c.db.Order("created_at ASC").Where("sender_wsp_number = ?", senderWspNumber).Limit(5).Find(&chatHistory)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetChatHistory] Error fetching chat history")
		return nil, errors.New("error fetching chat history")
	}

	return chatHistory, nil
}

// SaveChatHistory implements ChatHistoryRepository.
func (c *ChatHistoryRepositoryImpl) SaveChatHistory(senderWspNumber string, message string, botResponse string) error {

	chatHistory := models.ChatHistory{
		SenderWspNumber: senderWspNumber,
		Message:         message,
		BotResponse:     botResponse,
	}

	result := c.db.Create(&chatHistory)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [SaveChatHistory] Error saving chat history")
		return errors.New("error saving chat history")
	}

	return nil
}

func NewChatHistoryRepositoryImpl(db *gorm.DB) ChatHistoryRepository {
	return &ChatHistoryRepositoryImpl{db: db}
}
