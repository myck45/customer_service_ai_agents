package controllers

import "github.com/gin-gonic/gin"

type MenuController interface {
	CreateMenu(c *gin.Context)
	GetMenuByID(c *gin.Context)
	GetAllMenus(c *gin.Context)
	SemanticSearchMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
}
