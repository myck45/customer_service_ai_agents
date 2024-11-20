package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/service"
	"github.com/proyectos01-a/shared/handlers"
)

type MenuControllerImpl struct {
	menuService     service.MenuService
	responseHandler handlers.ResponseHandlers
}

// CreateMenu implements MenuController.
func (m *MenuControllerImpl) CreateMenu(c *gin.Context) {

	createMenuReq := &req.CreateMenuReq{}
	if err := c.ShouldBindJSON(createMenuReq); err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	if err := m.menuService.CreateMenu(createMenuReq); err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error creating menu", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu created successfully", nil)
}

// DeleteMenu implements MenuController.
func (m *MenuControllerImpl) DeleteMenu(c *gin.Context) {

	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := m.menuService.DeleteMenu(uint(id)); err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error deleting menu", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu deleted successfully", nil)
}

// GetAllMenus implements MenuController.
func (m *MenuControllerImpl) GetAllMenus(c *gin.Context) {

	menus, err := m.menuService.GetAllMenus()
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching menus", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menus fetched successfully", menus)
}

// GetMenuByID implements MenuController.
func (m *MenuControllerImpl) GetMenuByID(c *gin.Context) {
	panic("unimplemented")
}

// SemanticSearchMenu implements MenuController.
func (m *MenuControllerImpl) SemanticSearchMenu(c *gin.Context) {

	semanticSearchReq := &req.SemanticSearchReq{}
	if err := c.ShouldBindJSON(semanticSearchReq); err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	menus, err := m.menuService.SemanticSearchMenu(semanticSearchReq)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching menus", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menus fetched successfully", menus)
}

// UpdateMenu implements MenuController.
func (m *MenuControllerImpl) UpdateMenu(c *gin.Context) {

	updateMenuReq := &req.UpdateMenuReq{}
	if err := c.ShouldBindJSON(updateMenuReq); err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	menu, err := m.menuService.UpdateMenu(uint(id), updateMenuReq)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error updating menu", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu updated successfully", menu.ID)
}

func NewMenuControllerImpl(menuService service.MenuService, responseHandler handlers.ResponseHandlers) MenuController {
	return &MenuControllerImpl{
		menuService:     menuService,
		responseHandler: responseHandler,
	}
}
