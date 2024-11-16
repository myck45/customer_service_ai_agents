package data

import "github.com/proyectos01-a/RestaurantMenu/src/models"

type BotRepository interface {
	GetBotByWspNumber(wspNumber string) (*models.Bot, error)
	GetBotByID(id uint) (*models.Bot, error)
	GetBotByRestaurantID(restaurantID uint) ([]models.Bot, error)
	GetAllBots() ([]models.Bot, error)
	SaveBot(bot *models.Bot) error
	UpdateBot(bot *models.Bot) error
	DeleteBotByID(id uint) error
}
