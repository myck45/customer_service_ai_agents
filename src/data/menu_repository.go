package data

import (
	"github.com/pgvector/pgvector-go"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
)

type MenuRepository interface {
	CreateMenu(menu *models.Menu) error
	GetMenuByID(id uint) (*models.Menu, error)
	GetAllMenus() ([]models.Menu, error)
	SemanticSearchMenu(queryEmbedding pgvector.Vector, similarityThreshold float32, matchCount int) ([]models.Menu, error)
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id uint) error
}
