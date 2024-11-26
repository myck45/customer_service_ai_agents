package data

import (
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/models"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetMenuByID(id uint) (*models.Menu, error)
	GetAllMenus() ([]models.Menu, error)
	SemanticSearchMenu(queryEmbedding []float32, similarityThreshold float32, matchCount int, restaurantID uint) ([]dto.MenuSearchResponse, error)
	SemanticSearchWithSupabase(queryEmbedding []float32, similarityThreshold float32, matchCount int, restaurantID uint) ([]dto.MenuSearchResponse, error)
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id uint) error
}
