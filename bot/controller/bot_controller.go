package controller

import "github.com/gin-gonic/gin"

type BotController interface {
	BotResponse(c *gin.Context)
}
