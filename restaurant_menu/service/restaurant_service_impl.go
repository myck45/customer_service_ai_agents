package service

import (
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
)

type RestaurantServiceImpl struct {
	restaurantRepo data.RestaurantRepository
}

// CreateRestaurant implements RestaurantService.
func (r *RestaurantServiceImpl) CreateRestaurant(req *req.CreateRestaurantReq) (*res.RestaurantResponse, error) {

	restaurant := &models.Restaurant{
		Name: req.Name,
	}

	restaurant, err := r.restaurantRepo.CreateRestaurant(restaurant)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error creating restaurant")
		return nil, err
	}

	return &res.RestaurantResponse{
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
func (r *RestaurantServiceImpl) GetAllRestaurants() ([]res.RestaurantResponse, error) {
	restaurants, err := r.restaurantRepo.GetAllRestaurants()
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error fetching restaurants")
		return nil, err
	}

	restaurantList := make([]res.RestaurantResponse, 0, len(restaurants))
	for _, restaurant := range restaurants {
		restaurantList = append(restaurantList, res.RestaurantResponse{
			ID:   restaurant.ID,
			Name: restaurant.Name,
		})
	}

	return restaurantList, nil
}

// GetRestaurantByID implements RestaurantService.
func (r *RestaurantServiceImpl) GetRestaurantByID(id uint) (*res.RestaurantResponse, error) {
	restaurant, err := r.restaurantRepo.GetRestaurantByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [RestaurantServiceImpl] Error fetching restaurant")
		return nil, err
	}

	return &res.RestaurantResponse{
		ID:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

// UpdateRestaurant implements RestaurantService.
func (r *RestaurantServiceImpl) UpdateRestaurant(id uint, req *req.UpdateRestaurantReq) (*res.RestaurantResponse, error) {
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

	return &res.RestaurantResponse{
		ID:   restaurant.ID,
		Name: restaurant.Name,
	}, nil
}

func NewRestaurantServiceImpl(restaurantRepo data.RestaurantRepository) RestaurantService {
	return &RestaurantServiceImpl{
		restaurantRepo: restaurantRepo,
	}
}
