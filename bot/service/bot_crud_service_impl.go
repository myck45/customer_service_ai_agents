package service

import (
	"fmt"

	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/bot/dto/res"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
)

type BotCRUDServiceImpl struct {
	botRepo data.BotRepository
}

// CreateBot implements BotCRUDService.
func (b *BotCRUDServiceImpl) CreateBot(bot req.CreateBotReq) (*res.BotResponse, error) {
	// Create the bot
	newBot := &models.Bot{
		Name:         bot.Name,
		Identity:     bot.Identity,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	// Save the bot
	if err := b.botRepo.SaveBot(newBot); err != nil {
		logrus.WithError(err).Error("failed to save bot")
		return nil, fmt.Errorf("failed to save bot: %v", err)
	}

	// Create the response
	botResponse := &res.BotResponse{
		ID:           newBot.ID,
		Name:         newBot.Name,
		Identity:     newBot.Identity,
		WspNumber:    newBot.WspNumber,
		RestaurantID: newBot.RestaurantID,
	}

	return botResponse, nil
}

// DeleteBotByID implements BotCRUDService.
func (b *BotCRUDServiceImpl) DeleteBotByID(botID uint) error {
	// Delete the bot
	if err := b.botRepo.DeleteBotByID(botID); err != nil {
		logrus.WithError(err).Error("failed to delete bot")
		return fmt.Errorf("failed to delete bot: %v", err)
	}

	return nil
}

// GetAllBots implements BotCRUDService.
func (b *BotCRUDServiceImpl) GetAllBots() ([]res.BotResponse, error) {

	// Get all bots
	bots, err := b.botRepo.GetAllBots()
	if err != nil {
		logrus.WithError(err).Error("failed to get all bots")
		return nil, fmt.Errorf("failed to get all bots: %v", err)
	}

	// Create the response
	var botResponses []res.BotResponse
	for _, bot := range bots {
		botResponses = append(botResponses, res.BotResponse{
			ID:           bot.ID,
			Name:         bot.Name,
			Identity:     bot.Identity,
			WspNumber:    bot.WspNumber,
			RestaurantID: bot.RestaurantID,
		})
	}

	return botResponses, nil
}

// GetBotByID implements BotCRUDService.
func (b *BotCRUDServiceImpl) GetBotByID(botID uint) (*res.BotResponse, error) {

	// Get the bot
	bot, err := b.botRepo.GetBotByID(botID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	botResponse := &res.BotResponse{
		ID:           bot.ID,
		Name:         bot.Name,
		Identity:     bot.Identity,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	return botResponse, nil
}

// GetBotByRestaurantID implements BotCRUDService.
func (b *BotCRUDServiceImpl) GetBotByRestaurantID(restaurantID uint) ([]res.BotResponse, error) {

	// Get the bots
	bots, err := b.botRepo.GetBotByRestaurantID(restaurantID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	var botResponses []res.BotResponse
	for _, bot := range bots {
		botResponses = append(botResponses, res.BotResponse{
			ID:           bot.ID,
			Name:         bot.Name,
			Identity:     bot.Identity,
			WspNumber:    bot.WspNumber,
			RestaurantID: bot.RestaurantID,
		})
	}

	return botResponses, nil
}

// GetBotByWspNumber implements BotCRUDService.
func (b *BotCRUDServiceImpl) GetBotByWspNumber(wspNumber string) (*res.BotResponse, error) {

	// Get the bot
	bot, err := b.botRepo.GetBotByWspNumber(wspNumber)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Create the response
	botResponse := &res.BotResponse{
		ID:           bot.ID,
		Name:         bot.Name,
		Identity:     bot.Identity,
		WspNumber:    bot.WspNumber,
		RestaurantID: bot.RestaurantID,
	}

	return botResponse, nil
}

// UpdateBot implements BotCRUDService.
func (b *BotCRUDServiceImpl) UpdateBot(botID uint, bot req.UpdateBotReq) (*res.BotResponse, error) {

	// Get the bot
	botToUpdate, err := b.botRepo.GetBotByID(botID)
	if err != nil {
		logrus.WithError(err).Error("failed to get bot")
		return nil, fmt.Errorf("failed to get bot: %v", err)
	}

	// Update the bot
	botToUpdate.Name = bot.Name
	botToUpdate.Identity = bot.Identity
	botToUpdate.WspNumber = bot.WspNumber

	// Save the bot
	if err := b.botRepo.SaveBot(botToUpdate); err != nil {
		logrus.WithError(err).Error("failed to save bot")
		return nil, fmt.Errorf("failed to save bot: %v", err)
	}

	// Create the response
	botResponse := &res.BotResponse{
		ID:           botToUpdate.ID,
		Name:         botToUpdate.Name,
		WspNumber:    botToUpdate.WspNumber,
		RestaurantID: botToUpdate.RestaurantID,
	}

	return botResponse, nil
}

func NewBotCRUDServiceImpl(botRepo data.BotRepository) BotCRUDService {
	return &BotCRUDServiceImpl{
		botRepo: botRepo,
	}
}
