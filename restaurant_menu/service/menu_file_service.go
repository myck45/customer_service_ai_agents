package service

import (
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
)

type MenuFileService interface {
	// CreateMenuFile creates a new menu file
	CreateMenuFile(fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error)

	// GetMenuFileByRestaurantID gets all menu files by restaurant ID
	GetMenuFileByRestaurantID(restaurantID uint) ([]res.MenuFileResponse, error)

	// DeleteMenuFile deletes a menu file by ID
	DeleteMenuFile(fileID uint) error

	// GetMenuFileByID gets a menu file by ID
	GetMenuFileByID(fileID uint) (*res.MenuFileResponse, error)

	// GetMenuFilesURLByRestaurantID gets all menu files URLs by restaurant ID
	GetMenuFilesURLByRestaurantID(restaurantID uint) ([]string, error)

	// UpdateMenuFile updates a menu file by ID
	UpdateMenuFile(fileID uint, fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error)
}
