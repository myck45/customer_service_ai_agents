package service

import (
	"encoding/json"
	"errors"

	"github.com/pgvector/pgvector-go"
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/models"
	"github.com/proyectos01-a/shared/utils"
	"github.com/sirupsen/logrus"
)

type MenuServiceImpl struct {
	menuRepo data.MenuRepository
	botUtils utils.BotUtils
}

// CreateMenu implements MenuService.
func (m *MenuServiceImpl) CreateMenu(req *req.CreateMenuReq) error {
	if req == nil {
		logrus.Error("*** [CreateMenu] Invalid null request")
		return errors.New("invalid null request")
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		logrus.WithError(err).Error("*** [CreateMenu] Error marshalling request")
		return err
	}

	reqString := string(reqJSON)

	embedding, err := m.botUtils.GenerateEmbedding(reqString)
	if err != nil {
		logrus.WithError(err).Error("*** [CreateMenu] Error generating embedding")
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
		logrus.WithError(err).Error("*** [CreateMenu] Error creating menu")
		return err
	}

	return nil
}

// DeleteMenu implements MenuService.
func (m *MenuServiceImpl) DeleteMenu(id uint) error {
	err := m.menuRepo.DeleteMenu(id)
	if err != nil {
		logrus.WithError(err).Error("*** [DeleteMenu] Error deleting menu")
		return err
	}

	return nil
}

// GetAllMenus implements MenuService.
func (m *MenuServiceImpl) GetAllMenus() ([]res.MenuResponse, error) {

	menuList, err := m.menuRepo.GetAllMenus()
	if err != nil {
		logrus.WithError(err).Error("*** [GetAllMenus] Error getting all menus")
		return nil, err
	}

	var resList []res.MenuResponse
	for _, menu := range menuList {
		resList = append(resList, res.MenuResponse{
			ID:           menu.ID,
			RestaurantID: menu.RestaurantID,
			ItemName:     menu.ItemName,
			Description:  menu.Description,
			Price:        menu.Price,
			Likes:        menu.Likes,
		})
	}

	return resList, nil
}

// GetMenuByID implements MenuService.
func (m *MenuServiceImpl) GetMenuByID(id uint) (*res.MenuResponse, error) {

	menu, err := m.menuRepo.GetMenuByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [GetMenuByID] Error fetching menu")
		return nil, err
	}

	return &res.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		ItemName:     menu.ItemName,
		Description:  menu.Description,
		Price:        menu.Price,
		Likes:        menu.Likes,
	}, nil
}

// SemanticSearchMenu implements MenuService.
func (m *MenuServiceImpl) SemanticSearchMenu(req *req.SemanticSearchReq) ([]dto.MenuSearchResponse, error) {
	if req == nil {
		logrus.Error("*** [SemanticSearchMenu] Invalid null request")
		return nil, errors.New("invalid null request")
	}

	queryEmbedding, err := m.botUtils.GenerateEmbedding(req.Query)
	if err != nil {
		logrus.WithError(err).Error("*** [SemanticSearchMenu] Error generating embedding")
		return nil, err
	}

	menus, err := m.menuRepo.SemanticSearchMenu(queryEmbedding, req.SimilarityThreshold, req.MatchCount, req.RestaurantID)
	if err != nil {
		logrus.WithError(err).Error("*** [SemanticSearchMenu] Error fetching menus")
		return nil, err
	}

	menuList := make([]dto.MenuSearchResponse, 0, len(menus))
	for _, menu := range menus {
		menuList = append(menuList, dto.MenuSearchResponse{
			ID:          menu.ID,
			ItemName:    menu.ItemName,
			Description: menu.Description,
			Price:       menu.Price,
			Likes:       menu.Likes,
		})
	}

	return menuList, nil
}

// UpdateMenu implements MenuService.
func (m *MenuServiceImpl) UpdateMenu(id uint, req *req.UpdateMenuReq) (*res.MenuResponse, error) {

	if req == nil {
		logrus.Error("*** [UpdateMenu] Invalid null request")
		return nil, errors.New("invalid null request")
	}

	menu, err := m.menuRepo.GetMenuByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [UpdateMenu] Error fetching menu")
		return nil, err
	}

	menu.ItemName = req.ItemName
	menu.Description = req.Description
	menu.Price = req.Price
	menu.Likes = req.Likes

	if err = m.menuRepo.UpdateMenu(menu); err != nil {
		logrus.WithError(err).Error("*** [UpdateMenu] Error updating menu")
		return nil, err
	}

	return &res.MenuResponse{
		ID:           menu.ID,
		RestaurantID: menu.RestaurantID,
		ItemName:     menu.ItemName,
		Description:  menu.Description,
		Price:        menu.Price,
		Likes:        menu.Likes,
	}, nil

}

func NewMenuServiceImpl(menuRepo data.MenuRepository, botUtils utils.BotUtils) MenuService {
	return &MenuServiceImpl{
		menuRepo: menuRepo,
		botUtils: botUtils,
	}
}
