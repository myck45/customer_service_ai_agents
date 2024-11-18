package controller

import "github.com/gin-gonic/gin"

type RestaurantController interface {
	CreateRestaurant(c *gin.Context)
	GetRestaurantByID(c *gin.Context)
	GetAllRestaurants(c *gin.Context)
	UpdateRestaurant(c *gin.Context)
	DeleteRestaurant(c *gin.Context)
}
