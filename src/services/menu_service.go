package services

import (
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
)

type MenuService interface {
	// Create
	CreateMenu(req *request.CreateMenuReq) error

	// Read operations
	GetMenuByID(id uint) (*response.MenuResponse, error)
	GetAllMenus() (*response.MenuListResponse, error)

	// Semantic search
	SemanticSearchMenu(req *request.SemanticSearchReq) ([]response.MenuSearchResponse, error)

	// Update
	UpdateMenu(id uint, req *request.UpdateMenuReq) (*response.MenuResponse, error)

	// Delete
	DeleteMenu(id uint) error
}
