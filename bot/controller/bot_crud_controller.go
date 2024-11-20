package controller

import "github.com/gin-gonic/gin"

type BotCRUDController interface {
	GetBotByID(c *gin.Context)
	GetBotByRestaurantID(c *gin.Context)
	GetBotByWspNumber(c *gin.Context)
	GetAllBots(c *gin.Context)
	CreateBot(c *gin.Context)
	UpdateBot(c *gin.Context)
	DeleteBotByID(c *gin.Context)
}
