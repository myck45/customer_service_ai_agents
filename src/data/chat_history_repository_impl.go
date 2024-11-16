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
func (c *ChatHistoryRepositoryImpl) GetChatHistory(senderWspNumber string, botWspNumber string, restaurantID uint) ([]models.ChatHistory, error) {

	var chatHistory []models.ChatHistory

	result := c.db.
		Order("created_at ASC").
		Where("sender_wsp_number = ?", senderWspNumber).
		Where("bot_wsp_number = ?", botWspNumber).
		Where("restaurant_id = ?", restaurantID).
		Where("DATE(created_at) = DATE(NOW())").
		Limit(5).
		Find(&chatHistory)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetChatHistory] Error fetching chat history")
		return nil, errors.New("error fetching chat history")
	}

	return chatHistory, nil
}

// GetChatHistoryBySenderWspNumberAndRestaurantID implements ChatHistoryRepository.
func (c *ChatHistoryRepositoryImpl) GetChatHistoryBySenderWspNumberAndRestaurantID(senderWspNumber string, restaurantID uint) ([]models.ChatHistory, error) {

	var chatHistory []models.ChatHistory

	result := c.db.
		Order("created_at ASC").
		Where("sender_wsp_number = ?", senderWspNumber).
		Where("restaurant_id = ?", restaurantID).
		Where("DATE(created_at) = DATE(NOW())").
		Limit(5).
		Find(&chatHistory)

	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetChatHistoryBySenderWspNumberAndRestaurantID] Error fetching chat history")
		return nil, errors.New("error fetching chat history")
	}

	return chatHistory, nil
}

// SaveChatHistory implements ChatHistoryRepository.
func (c *ChatHistoryRepositoryImpl) SaveChat(chatHistory *models.ChatHistory) error {

	result := c.db.Create(chatHistory)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [SaveChatHistory] Error saving chat history")
		return errors.New("error saving chat history")
	}

	return nil
}

func NewChatHistoryRepositoryImpl(db *gorm.DB) ChatHistoryRepository {
	return &ChatHistoryRepositoryImpl{db: db}
}
