package data

import "github.com/proyectos01-a/shared/models"

type MenuFileRepository interface {
	SaveMenuFile(menuFile *models.MenuFile) (*models.MenuFile, error)
	GetMenuFileByRestaurantID(restaurantID uint) ([]models.MenuFile, error)
	GetMenuFileByID(menuFileID uint) (*models.MenuFile, error)
	DeleteMenuFile(menuFileID uint) error
	UpdateMenuFile(menuFile *models.MenuFile) (*models.MenuFile, error)
}
