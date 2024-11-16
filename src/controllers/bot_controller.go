package controllers

import "github.com/gin-gonic/gin"

type BotController interface {
	HandleTwilioWebhook(c *gin.Context)
	GetBotByID(c *gin.Context)
	GetBotByRestaurantID(c *gin.Context)
	GetBotByWspNumber(c *gin.Context)
	GetAllBots(c *gin.Context)
	CreateBot(c *gin.Context)
	UpdateBot(c *gin.Context)
	DeleteBotByID(c *gin.Context)
}
