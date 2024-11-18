package data

import (
	"fmt"

	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RestaurantRepositoryImpl struct {
	db *gorm.DB
}

// CreateRestaurant implements RestaurantRepository.
func (r *RestaurantRepositoryImpl) CreateRestaurant(restaurant *models.Restaurant) (*models.Restaurant, error) {
	result := r.db.Create(restaurant)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error creating restaurant")
		return nil, fmt.Errorf("error creating restaurant")
	}

	return restaurant, nil
}

// DeleteRestaurant implements RestaurantRepository.
func (r *RestaurantRepositoryImpl) DeleteRestaurant(id uint) error {
	result := r.db.Delete(&models.Restaurant{}, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error deleting restaurant")
		return fmt.Errorf("error deleting restaurant with id %d", id)
	}

	if result.RowsAffected == 0 {
		logrus.WithField("id", id).Warn("Restaurant not found")
		return fmt.Errorf("restaurant with id %d not found", id)
	}

	return nil
}

// GetAllRestaurants implements RestaurantRepository.
func (r *RestaurantRepositoryImpl) GetAllRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	result := r.db.Find(&restaurants)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error fetching restaurants")
		return nil, fmt.Errorf("error fetching restaurants")
	}

	return restaurants, nil
}

// GetRestaurantByID implements RestaurantRepository.
func (r *RestaurantRepositoryImpl) GetRestaurantByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	result := r.db.First(&restaurant, id)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error fetching restaurant")
		return nil, fmt.Errorf("error fetching restaurant with id %d", id)
	}

	return &restaurant, nil
}

// UpdateRestaurant implements RestaurantRepository.
func (r *RestaurantRepositoryImpl) UpdateRestaurant(restaurant *models.Restaurant) error {
	result := r.db.Save(restaurant)
	if result.Error != nil {
		logrus.WithError(result.Error).Error("Error updating restaurant")
		return fmt.Errorf("error updating restaurant with id %d", restaurant.ID)
	}

	return nil
}

func NewRestaurantRepositoryImpl(db *gorm.DB) RestaurantRepository {
	return &RestaurantRepositoryImpl{
		db: db,
	}
}
