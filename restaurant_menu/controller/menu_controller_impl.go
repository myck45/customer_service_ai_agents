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

// CreateMenu godoc
//
//	@Summary		Create a new menu
//	@Description	create a new menu with the input payload
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Param			request			body		req.CreateMenuReq	true	"Create Menu Request"
//	@Success		200				{object}	dto.BaseResponse	"Menu created successfully"
//	@Failure		400				{object}	dto.BaseResponse	"Error binding request"
//	@Failure		500				{object}	dto.BaseResponse	"Error creating menu"
//	@Router			/api/v1/menu	[post]
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

// DeleteMenu godoc
//
//	@Summary		Delete a menu
//	@Description	delete a menu with the input id
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int					true	"Menu ID"
//	@Success		200					{object}	dto.BaseResponse	"Menu deleted successfully"
//	@Failure		400					{object}	dto.BaseResponse	"Error parsing id"
//	@Failure		500					{object}	dto.BaseResponse	"Error deleting menu"
//	@Router			/api/v1/menu/{id}	[delete]
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

// GetAllMenus godoc
//
//	@Summary		Get all menus
//	@Description	get all menus
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	dto.BaseResponse{data=[]res.MenuResponse}	"Menus fetched successfully"
//	@Failure		500				{object}	dto.BaseResponse							"Error fetching menus"
//	@Router			/api/v1/menu	[get]
func (m *MenuControllerImpl) GetAllMenus(c *gin.Context) {

	menus, err := m.menuService.GetAllMenus()
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching menus", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menus fetched successfully", menus)
}

// GetMenuByID godoc
//
//	@Summary		Get a menu by ID
//	@Description	get a menu by the input id
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int										true	"Menu ID"
//	@Success		200					{object}	dto.BaseResponse{data=res.MenuResponse}	"Menu fetched successfully"
//	@Failure		400					{object}	dto.BaseResponse						"Error parsing id"
//	@Failure		500					{object}	dto.BaseResponse						"Error fetching menu"
//	@Router			/api/v1/menu/{id}	[get]
func (m *MenuControllerImpl) GetMenuByID(c *gin.Context) {

	menuID := c.Param("id")
	id, err := strconv.ParseUint(menuID, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	menu, err := m.menuService.GetMenuByID(uint(id))
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching menu", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu fetched successfully", menu)
}

// SemanticSearchMenu godoc
//
//	@Summary		Semantic search menu
//	@Description	semantic search menu with the input payload
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Param			request				body		req.SemanticSearchReq							true	"Semantic Search Request"
//	@Success		200					{object}	dto.BaseResponse{data=[]dto.MenuSearchResponse}	"Menus fetched successfully"
//	@Failure		400					{object}	dto.BaseResponse								"Error binding request"
//	@Failure		500					{object}	dto.BaseResponse								"Error fetching menus"
//	@Router			/api/v1/menu/search	[get]
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

// UpdateMenu godoc
//
//	@Summary		Update a menu
//	@Description	update a menu with the input id and payload
//	@Tags			menu
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int					true	"Menu ID"
//	@Param			request				body		req.UpdateMenuReq	true	"Update Menu Request"
//	@Success		200					{object}	dto.BaseResponse	"Menu updated successfully"
//	@Failure		400					{object}	dto.BaseResponse	"Error binding request"
//	@Failure		500					{object}	dto.BaseResponse	"Error updating menu"
//	@Router			/api/v1/menu/{id}	[put]
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
