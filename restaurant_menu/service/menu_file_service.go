package service

import (
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
)

type MenuFileService interface {
	CreateMenuFile(fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error)
	GetMenuFileByRestaurantID(restaurantID uint) ([]res.MenuFileResponse, error)
	DeleteMenuFile(fileID uint) error
	GetMenuFileByID(fileID uint) (*res.MenuFileResponse, error)
	UpdateMenuFile(fileID uint, fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error)
}
