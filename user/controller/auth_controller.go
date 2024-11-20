package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(c *gin.Context)
}
