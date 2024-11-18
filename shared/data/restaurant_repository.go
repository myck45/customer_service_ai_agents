package data

import "github.com/proyectos01-a/shared/models"

type RestaurantRepository interface {
	CreateRestaurant(restaurant *models.Restaurant) (*models.Restaurant, error)
	GetRestaurantByID(id uint) (*models.Restaurant, error)
	GetAllRestaurants() ([]models.Restaurant, error)
	UpdateRestaurant(restaurant *models.Restaurant) error
	DeleteRestaurant(id uint) error
}
