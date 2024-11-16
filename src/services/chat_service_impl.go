package services

import (
	"github.com/proyectos01-a/RestaurantMenu/src/data"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sirupsen/logrus"
)

type ChatServiceImpl struct {
	chatRepo data.ChatHistoryRepository
	botRepo  data.BotRepository
}

// GetChatHistoryBySenderWspNumber implements ChatService.
func (c *ChatServiceImpl) GetChatHistoryBySenderWspNumber(senderWspNumber string) ([]models.ChatHistory, error) {

	bot, err := c.botRepo.GetBotByWspNumber(senderWspNumber)
	if err != nil {
		logrus.WithError(err).Error("*** [GetChatHistoryBySenderWspNumber] Error getting bot")
		return nil, err
	}

	restaurantID := bot.RestaurantID

	chatHistory, err := c.chatRepo.GetChatHistoryBySenderWspNumberAndRestaurantID(senderWspNumber, restaurantID)
	if err != nil {
		logrus.WithError(err).Error("*** [GetChatHistoryBySenderWspNumber] Error getting chat history")
		return nil, err
	}

	return chatHistory, nil
}

// SaveChat implements ChatService.
func (c *ChatServiceImpl) SaveChat(chat request.TwilioWebhook, botResponse string) error {

	bot, err := c.botRepo.GetBotByWspNumber(chat.To)
	if err != nil {
		logrus.WithError(err).Error("*** [SaveChat] Error getting bot")
		return err
	}

	chatHistory := &models.ChatHistory{
		SenderWspNumber: chat.From,
		Message:         chat.Body,
		BotResponse:     botResponse,
		RestaurantID:    bot.RestaurantID,
	}

	err = c.chatRepo.SaveChat(chatHistory)
	if err != nil {
		logrus.WithError(err).Error("*** [SaveChat] Error saving chat")
		return err
	}

	return nil
}

func NewChatServiceImpl(chatRepo data.ChatHistoryRepository, botRepo data.BotRepository) ChatService {
	return &ChatServiceImpl{
		chatRepo: chatRepo,
		botRepo:  botRepo,
	}
}
