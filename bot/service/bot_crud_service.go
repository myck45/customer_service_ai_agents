package service

import (
	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/bot/dto/res"
)

type BotCRUDService interface {
	GetBotByID(botID uint) (*res.BotResponse, error)
	GetBotByRestaurantID(restaurantID uint) ([]res.BotResponse, error)
	GetBotByWspNumber(wspNumber string) (*res.BotResponse, error)
	GetAllBots() ([]res.BotResponse, error)
	CreateBot(bot *req.CreateBotReq) (*res.BotResponse, error)
	UpdateBot(botID uint, bot *req.UpdateBotReq) (*res.BotResponse, error)
	DeleteBotByID(botID uint) error
}
