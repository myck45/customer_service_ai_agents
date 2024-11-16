package data

import (
	"errors"

	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BotRepositoryImpl struct {
	db *gorm.DB
}

// DeleteBot implements BotRepository.
func (b *BotRepositoryImpl) DeleteBotByID(id uint) error {

	result := b.db.Delete(&models.Bot{}, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [DeleteBotByID] Error deleting bot")
		return errors.New("error deleting bot")
	}

	if result.RowsAffected == 0 {
		logrus.WithField("id", id).Warn("*** [DeleteBotByID] Bot not found")
		return errors.New("bot not found")
	}

	return nil
}

// GetAllBots implements BotRepository.
func (b *BotRepositoryImpl) GetAllBots() ([]models.Bot, error) {
	var bots []models.Bot

	result := b.db.Find(&bots)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetAllBots] Error fetching bots")
		return nil, errors.New("error fetching bots")
	}

	return bots, nil
}

// GetBotByID implements BotRepository.
func (b *BotRepositoryImpl) GetBotByID(id uint) (*models.Bot, error) {
	var bot models.Bot

	result := b.db.First(&bot, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetBotByID] Error fetching bot")
		return nil, errors.New("error fetching bot")
	}

	return &bot, nil
}

// GetBotByRestaurantID implements BotRepository.
func (b *BotRepositoryImpl) GetBotByRestaurantID(restaurantID uint) ([]models.Bot, error) {
	var bots []models.Bot

	result := b.db.Where("restaurant_id = ?", restaurantID).Find(&bots)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetBotByRestaurantID] Error fetching bots")
		return nil, errors.New("error fetching bots")
	}

	return bots, nil
}

// GetBotByWspNumber implements BotRepository.
func (b *BotRepositoryImpl) GetBotByWspNumber(wspNumber string) (*models.Bot, error) {

	var bot models.Bot

	result := b.db.Where("wsp_number = ?", wspNumber).First(&bot)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [GetBotByWspNumber] Error fetching bot")
		return nil, errors.New("error fetching bot")
	}

	return &bot, nil
}

// SaveBot implements BotRepository.
func (b *BotRepositoryImpl) SaveBot(bot *models.Bot) error {
	result := b.db.Create(bot)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [SaveBot] Error creating bot")
		return errors.New("error creating bot")
	}

	return nil
}

// UpdateBot implements BotRepository.
func (b *BotRepositoryImpl) UpdateBot(bot *models.Bot) error {

	result := b.db.Save(bot)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("*** [UpdateBot] Error updating bot")
		return errors.New("error updating bot")
	}

	return nil
}

func NewBotRepositoryImpl(db *gorm.DB) BotRepository {
	return &BotRepositoryImpl{
		db: db,
	}
}
