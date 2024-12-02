package controller

import (
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/service"
	_ "github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/handlers"
)

type MenuFileControllerImpl struct {
	menuFileService service.MenuFileService
	responseHandler handlers.ResponseHandlers
}

// CreateMenuFile godoc
//
//	@Summary		Create a new menu file
//	@Description	create a new menu file with the input payload
//	@Tags			menu-file
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file				form		file				true	"Menu File"
//	@Param			restaurant_id		form		int					true	"Restaurant ID"
//	@Success		200					{object}	dto.BaseResponse	"Menu file created successfully"
//	@Failure		400					{object}	dto.BaseResponse	"Error getting file from request"
//	@Failure		413					{object}	dto.BaseResponse	"File exceeds maximum size of 16MB"
//	@Failure		500					{object}	dto.BaseResponse	"Error creating menu file"
//	@Router			/api/v1/menu-file	[post]
func (m *MenuFileControllerImpl) CreateMenuFile(c *gin.Context) {

	// Limit file size (16 MB)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 16<<20)

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error getting file from request", err)
		return
	}

	// Verify file size
	if file.Size > 16*1024*1024 {
		m.responseHandler.HandleError(c, http.StatusRequestEntityTooLarge, "File exceeds maximum size of 16MB", nil)
		return
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error opening file", err)
		return
	}
	defer src.Close()

	// Read file bytes
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error reading file", err)
		return
	}

	// Get restaurant ID from request
	restaurantIDStr := c.PostForm("restaurant_id")
	restaurantID, err := strconv.ParseUint(restaurantIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid restaurant ID", err)
		return
	}

	// Validate file extension
	allowedExt := []string{".pdf", ".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(file.Filename)
	if !m.ValidateMenuFileExt(allowedExt, strings.ToLower(ext)) {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid file extension", nil)
		return
	}

	// Prepare create menu file request
	createMenuFileReq := &req.CreateMenuFileReq{
		FileName:     file.Filename,
		FileType:     ext,
		RestaurantID: uint(restaurantID),
	}

	// Create menu file
	menuFile, err := m.menuFileService.CreateMenuFile(createMenuFileReq, fileBytes)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error creating menu file", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu file created successfully", menuFile)
}

// DeleteMenuFile godoc
//
//	@Summary		Delete a menu file
//	@Description	delete a menu file with the input id
//	@Tags			menu-file
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int					true	"Menu File ID"
//	@Success		200						{object}	dto.BaseResponse	"Menu file deleted successfully"
//	@Failure		400						{object}	dto.BaseResponse	"Invalid file ID"
//	@Failure		500						{object}	dto.BaseResponse	"Error deleting menu file"
//	@Router			/api/v1/menu-file/{id}	[delete]
func (m *MenuFileControllerImpl) DeleteMenuFile(c *gin.Context) {
	fileIDStr := c.Param("id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid file ID", err)
		return
	}

	if err := m.menuFileService.DeleteMenuFile(uint(fileID)); err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error deleting menu file", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu file deleted successfully", nil)
}

// GetMenuFileByID godoc
//
//	@Summary		Get a menu file by ID
//	@Description	get a menu file with the input id
//	@Tags			menu-file
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int											true	"Menu File ID"
//	@Success		200						{object}	dto.BaseResponse{data=res.MenuFileResponse}	"Menu file fetched successfully"
//	@Failure		400						{object}	dto.BaseResponse							"Invalid file ID"
//	@Failure		500						{object}	dto.BaseResponse							"Error getting menu file"
//	@Router			/api/v1/menu-file/{id}	[get]
func (m *MenuFileControllerImpl) GetMenuFileByID(c *gin.Context) {
	fileIDStr := c.Param("id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid file ID", err)
		return
	}

	menuFile, err := m.menuFileService.GetMenuFileByID(uint(fileID))
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error getting menu file", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu file fetched successfully", menuFile)
}

// GetMenuFileByRestaurantID godoc
//
//	@Summary		Get menu files by restaurant ID
//	@Description	get all menu files with the input restaurant ID
//	@Tags			menu-file
//	@Accept			json
//	@Produce		json
//	@Param			restaurant_id									path		int												true	"Restaurant ID"
//	@Success		200												{object}	dto.BaseResponse{data=[]res.MenuFileResponse}	"Menu files fetched successfully"
//	@Failure		400												{object}	dto.BaseResponse								"Invalid restaurant ID"
//	@Failure		500												{object}	dto.BaseResponse								"Error getting menu files"
//	@Router			/api/v1/menu-files/restaurant/{restaurant_id}	[get]
func (m *MenuFileControllerImpl) GetMenuFileByRestaurantID(c *gin.Context) {
	restaurantIDStr := c.Param("restaurant_id")
	restaurantID, err := strconv.ParseUint(restaurantIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid restaurant ID", err)
		return
	}

	menuFiles, err := m.menuFileService.GetMenuFileByRestaurantID(uint(restaurantID))
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error getting menu files", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu files fetched successfully", menuFiles)
}

// UpdateMenuFile godoc
//
//	@Summary		Update a menu file
//	@Description	update a menu file with the input payload
//	@Tags			menu-file
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id						path		int					true	"Menu File ID"
//	@Param			file					form		file				true	"Menu File"
//	@Param			restaurant_id			form		int					true	"Restaurant ID"
//	@Success		200						{object}	dto.BaseResponse	"Menu file updated successfully"
//	@Failure		400						{object}	dto.BaseResponse	"Invalid file ID"
//	@Failure		400						{object}	dto.BaseResponse	"Error getting file from request"
//	@Failure		500						{object}	dto.BaseResponse	"Error updating menu file"
//	@Router			/api/v1/menu-file/{id}	[put]
func (m *MenuFileControllerImpl) UpdateMenuFile(c *gin.Context) {
	fileIDStr := c.Param("id")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid file ID", err)
		return
	}

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error getting file from request", err)
		return
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error opening file", err)
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error reading file", err)
		return
	}

	// Validate file extension
	allowedExt := []string{".pdf", ".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(file.Filename)
	if !m.ValidateMenuFileExt(allowedExt, strings.ToLower(ext)) {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid file extension", nil)
		return
	}

	// Get restaurant ID from request
	restaurantIDStr := c.PostForm("restaurant_id")
	restaurantID, err := strconv.ParseUint(restaurantIDStr, 10, 64)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Invalid restaurant ID", err)
		return
	}

	// Prepare create menu file request
	createMenuFileReq := &req.CreateMenuFileReq{
		FileName:     file.Filename,
		FileType:     ext,
		RestaurantID: uint(restaurantID),
	}

	menuFile, err := m.menuFileService.UpdateMenuFile(uint(fileID), createMenuFileReq, fileBytes)
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusInternalServerError, "Error updating menu file", err)
		return
	}

	m.responseHandler.HandleSuccess(c, http.StatusOK, "Menu file updated successfully", menuFile)
}

// ValidateMenuFileExt implements MenuFileController.
func (m *MenuFileControllerImpl) ValidateMenuFileExt(ext []string, item string) bool {
	for _, e := range ext {
		if item == e {
			return true
		}
	}
	return false
}

func NewMenuFileControllerImpl(menuFileService service.MenuFileService, responseHandler handlers.ResponseHandlers) MenuFileController {
	return &MenuFileControllerImpl{
		menuFileService: menuFileService,
		responseHandler: responseHandler,
	}
}
