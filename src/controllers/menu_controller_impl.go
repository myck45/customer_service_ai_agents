package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/services"
)

type MenuControllerImpl struct {
	menuService      services.MenuService
	responseHandlers ResponseHandlers
}

// CreateMenu implements MenuController.
func (m *MenuControllerImpl) CreateMenu(c *gin.Context) {

	createMenuReq := &request.CreateMenuReq{}
	if err := c.ShouldBindJSON(createMenuReq); err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	if err := m.menuService.CreateMenu(createMenuReq); err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error creating menu", err)
		return
	}
	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menu created successfully", "Menu created successfully")
}

// DeleteMenu implements MenuController.
func (m *MenuControllerImpl) DeleteMenu(c *gin.Context) {
	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := m.menuService.DeleteMenu(uint(id)); err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error deleting menu", err)
		return
	}

	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menu deleted successfully", nil)
}

// GetAllMenus implements MenuController.
func (m *MenuControllerImpl) GetAllMenus(c *gin.Context) {
	menus, err := m.menuService.GetAllMenus()
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error fetching menus", err)
		return
	}

	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menus fetched successfully", menus)
}

// GetMenuByID implements MenuController.
func (m *MenuControllerImpl) GetMenuByID(c *gin.Context) {
	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	menu, err := m.menuService.GetMenuByID(uint(id))
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error fetching menu", err)
		return
	}

	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menu fetched successfully", menu)
}

// SemanticSearchMenu implements MenuController.
func (m *MenuControllerImpl) SemanticSearchMenu(c *gin.Context) {

	semanticSearchMenuReq := &request.SemanticSearchReq{}
	if err := c.ShouldBindJSON(semanticSearchMenuReq); err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	menus, err := m.menuService.SemanticSearchMenu(semanticSearchMenuReq)
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error fetching menus", err)
		return
	}

	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menus fetched successfully", menus)
}

// UpdateMenu implements MenuController.
func (m *MenuControllerImpl) UpdateMenu(c *gin.Context) {
	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	updateMenuReq := &request.UpdateMenuReq{}
	if err := c.ShouldBindJSON(updateMenuReq); err != nil {
		m.responseHandlers.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	menu, err := m.menuService.UpdateMenu(uint(id), updateMenuReq)
	if err != nil {
		m.responseHandlers.HandleError(c, http.StatusInternalServerError, "Error updating menu", err)
		return
	}

	m.responseHandlers.HandleSuccess(c, http.StatusOK, "Menu updated successfully", menu)
}

func NewMenuControllerImpl(menuService services.MenuService, responseHandlers ResponseHandlers) MenuController {
	return &MenuControllerImpl{
		menuService:      menuService,
		responseHandlers: responseHandlers,
	}
}
