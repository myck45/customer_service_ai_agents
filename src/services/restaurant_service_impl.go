package services

import (
	"github.com/proyectos01-a/RestaurantMenu/src/data"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sirupsen/logrus"
)

type RestaurantServiceImpl struct {
	restaurantRepo data.RestaurantRepository
}

// CreateRestaurant implements RestaurantService.
func (r *RestaurantServiceImpl) CreateRestaurant(req *request.CreateRestaurantReq) (*response.RestaurantResponse, error) {

	restaurant := &models.Restaurant{
		Name: req.Name,
	}

	restaurant, err := r.restaurantRepo.CreateRestaurant(restaurant)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error creating restaurant")
		return nil, err
	}

	return &response.RestaurantResponse{
		ID:   restaurant.ID,
		Name: restaurant.Name,
	}, nil

}

// DeleteRestaurant implements RestaurantService.
func (r *RestaurantServiceImpl) DeleteRestaurant(id uint) error {

	err := r.restaurantRepo.DeleteRestaurant(id)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error deleting restaurant")
		return err
	}

	return nil
}

// GetAllRestaurants implements RestaurantService.
func (r *RestaurantServiceImpl) GetAllRestaurants() (*response.RestaurantListResponse, error) {
	restaurants, err := r.restaurantRepo.GetAllRestaurants()
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error fetching restaurants")
		return nil, err
	}

	restaurantList := make([]response.RestaurantResponse, 0, len(restaurants))
	for _, restaurant := range restaurants {
		restaurantList = append(restaurantList, response.RestaurantResponse{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		})
	}

	return &response.RestaurantListResponse{
		Restaurants: restaurantList,
	}, nil
}

// GetRestaurantByID implements RestaurantService.
func (r *RestaurantServiceImpl) GetRestaurantByID(id uint) (*response.RestaurantResponse, error) {
	restaurant, err := r.restaurantRepo.GetRestaurantByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error fetching restaurant")
		return nil, err
	}

	return &response.RestaurantResponse{
		ID:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

// UpdateRestaurant implements RestaurantService.
func (r *RestaurantServiceImpl) UpdateRestaurant(id uint, req *request.UpdateRestaurantReq) (*response.RestaurantResponse, error) {
	restaurant, err := r.restaurantRepo.GetRestaurantByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error fetching restaurant")
		return nil, err
	}

	restaurant.Name = req.Name

	err = r.restaurantRepo.UpdateRestaurant(restaurant)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error updating restaurant")
		return nil, err
	}

	return &response.RestaurantResponse{
		ID:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func NewRestaurantServiceImpl(restaurantRepo data.RestaurantRepository) RestaurantService {
	return &RestaurantServiceImpl{
		restaurantRepo: restaurantRepo,
	}
}
