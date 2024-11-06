package services

import (
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
)

type RestaurantService interface {
	// Create - takes DTO, returns response
	CreateRestaurant(req *request.CreateRestaurantReq) (*response.RestaurantResponse, error)

	// Read - returns response DTOs
	GetRestaurantByID(id uint) (*response.RestaurantResponse, error)
	GetAllRestaurants() (*response.RestaurantListResponse, error)

	// Update - takes DTO, returns response
	UpdateRestaurant(id uint, req *request.UpdateRestaurantReq) (*response.RestaurantResponse, error)

	// Delete - just returns error
	DeleteRestaurant(id uint) error
}
