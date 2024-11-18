package service

import (
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
)

type RestaurantService interface {
	// Create - takes DTO, returns response
	CreateRestaurant(req *req.CreateRestaurantReq) (*res.RestaurantResponse, error)

	// Read - returns response DTOs
	GetRestaurantByID(id uint) (*res.RestaurantResponse, error)
	GetAllRestaurants() ([]res.RestaurantResponse, error)

	// Update - takes DTO, returns response
	UpdateRestaurant(id uint, req *req.UpdateRestaurantReq) (*res.RestaurantResponse, error)

	// Delete - just returns error
	DeleteRestaurant(id uint) error
}
