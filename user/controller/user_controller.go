package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
}
