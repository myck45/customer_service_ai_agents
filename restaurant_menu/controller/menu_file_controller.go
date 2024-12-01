package controller

import "github.com/gin-gonic/gin"

type MenuFileController interface {
	// CreateMenuFile creates a new menu file.
	CreateMenuFile(c *gin.Context)

	// DeleteMenuFile deletes a menu file.
	DeleteMenuFile(c *gin.Context)

	// GetMenuFileByID gets a menu file by its ID.
	GetMenuFileByID(c *gin.Context)

	// GetMenuFileByRestaurantID gets all menu files by restaurant ID.
	GetMenuFileByRestaurantID(c *gin.Context)

	// UpdateMenuFile updates a menu file.
	UpdateMenuFile(c *gin.Context)

	// ValidateMenuFileExt validates the file extension.
	ValidateMenuFileExt(ext []string, item string) bool
}
