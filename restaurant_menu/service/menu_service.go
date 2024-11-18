package service

import (
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
	"github.com/proyectos01-a/shared/dto"
)

type MenuService interface {
	// Create
	CreateMenu(req *req.CreateMenuReq) error

	// Read operations
	GetMenuByID(id uint) (*res.MenuResponse, error)
	GetAllMenus() ([]res.MenuResponse, error)

	// Semantic search
	SemanticSearchMenu(req *req.SemanticSearchReq) ([]dto.MenuSearchResponse, error)

	// Update
	UpdateMenu(id uint, req *req.UpdateMenuReq) (*res.MenuResponse, error)

	// Delete
	DeleteMenu(id uint) error
}
