package data

import (
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetMenuByID(id uint) (*models.Menu, error)
	GetAllMenus() ([]models.Menu, error)
	SemanticSearchMenu(queryEmbedding []float32, similarityThreshold float32, matchCount int, restaurantID uint) ([]response.MenuSearchResponse, error)
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id uint) error
}
