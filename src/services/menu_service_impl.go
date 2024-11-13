package services

import (
	"encoding/json"
	"errors"

	"github.com/pgvector/pgvector-go"
	"github.com/proyectos01-a/RestaurantMenu/src/data"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/proyectos01-a/RestaurantMenu/src/models"
	"github.com/sirupsen/logrus"
)

type MenuServiceImpl struct {
	bot      BotService
	menuRepo data.MenuRepository
}

// CreateMenu implements MenuService.
func (m *MenuServiceImpl) CreateMenu(req *request.CreateMenuReq) error {
	if req == nil {
		logrus.Error("Invalid null request")
		return errors.New("invalid null request")
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		logrus.WithError(err).Error("Error marshalling request")
		return err
	}

	reqString := string(reqJSON)

	embedding, err := m.bot.GenerateEmbedding(reqString)
	if err != nil {
		logrus.WithError(err).Error("Error generating embedding")
		return err
	}

	menuModel := &models.Menu{
		RestaurantID: req.RestaurantID,
		ItemName:     req.ItemName,
		Description:  req.Description,
		Price:        req.Price,
		Likes:        req.Likes,
		Embedding:    pgvector.NewVector(embedding),
	}

	err = m.menuRepo.CreateMenu(menuModel)
	if err != nil {
		logrus.WithError(err).Error("Error creating menu")
		return err
	}

	return nil

}

// DeleteMenu implements MenuService.
func (m *MenuServiceImpl) DeleteMenu(id uint) error {

	err := m.menuRepo.DeleteMenu(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllMenus implements MenuService.
func (m *MenuServiceImpl) GetAllMenus() (*response.MenuListResponse, error) {

	menus, err := m.menuRepo.GetAllMenus()
	if err != nil {
		return nil, err
	}

	menuList := make([]response.MenuResponse, 0, len(menus))
	for _, menu := range menus {
		menuList = append(menuList, response.MenuResponse{
			ID:           menu.ID,
			RestaurantID: menu.RestaurantID,
			ItemName:     menu.ItemName,
			Description:  menu.Description,
			Price:        menu.Price,
			Likes:        menu.Likes,
		})
	}

	return &response.MenuListResponse{
		Menus: menuList,
	}, nil
}

// GetMenuByID implements MenuService.
func (m *MenuServiceImpl) GetMenuByID(id uint) (*response.MenuResponse, error) {

	menu, err := m.menuRepo.GetMenuByID(id)
	if err != nil {
		return nil, err
	}

	return &response.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		ItemName:     menu.ItemName,
		Description:  menu.Description,
		Price:        menu.Price,
		Likes:        menu.Likes,
	}, nil
}

// SemanticSearchMenu implements MenuService.
func (m *MenuServiceImpl) SemanticSearchMenu(query string, similarityThreshold float32, matchCount int) (*response.MenuListResponse, error) {
	panic("unimplemented")
}

// UpdateMenu implements MenuService.
func (m *MenuServiceImpl) UpdateMenu(id uint, req *request.UpdateMenuReq) (*response.MenuResponse, error) {

	if req == nil {
		logrus.Error("Invalid null request")
		return nil, errors.New("invalid null request")
	}

	menu, err := m.menuRepo.GetMenuByID(id)
	if err != nil {
		logrus.WithError(err).Error("Error fetching menu")
		return nil, err
	}

	if req.ItemName != "" {
		menu.ItemName = req.ItemName
	} else {
		logrus.Error("Item name cannot be empty")
		return nil, errors.New("item name cannot be empty")
	}

	if req.Description != "" {
		menu.Description = req.Description
	} else {
		logrus.Error("Description cannot be empty")
		return nil, errors.New("description cannot be empty")
	}

	if req.Price != 0 {
		menu.Price = req.Price
	} else {
		logrus.Error("Price cannot be 0")
		return nil, errors.New("price cannot be 0")
	}

	if req.Likes != 0 {
		menu.Likes = req.Likes
	} else {
		logrus.Error("Likes cannot be 0")
		return nil, errors.New("likes cannot be 0")
	}

	err = m.menuRepo.UpdateMenu(menu)
	if err != nil {
		logrus.WithError(err).Error("Error updating menu")
		return nil, err
	}

	return &response.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		ItemName:     menu.ItemName,
		Description:  menu.Description,
		Price:        menu.Price,
		Likes:        menu.Likes,
	}, nil
}

func NewMenuServiceImpl(bot BotService, menuRepo data.MenuRepository) MenuService {
	return &MenuServiceImpl{
		bot:      bot,
		menuRepo: menuRepo,
	}
}
