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
	"github.com/proyectos01-a/shared/handlers"
)

type MenuFileControllerImpl struct {
	menuFileService service.MenuFileService
	responseHandler handlers.ResponseHandlers
}

// CreateMenuFile implements MenuFileController.
func (m *MenuFileControllerImpl) CreateMenuFile(c *gin.Context) {

	// Limit file size (16 MB)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 16<<20)

	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		m.responseHandler.HandleError(c, http.StatusBadRequest, "Error getting file from request", err)
		return
	}

	// Verificación adicional de tamaño
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

// DeleteMenuFile implements MenuFileController.
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

// GetMenuFileByID implements MenuFileController.
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

// GetMenuFileByRestaurantID implements MenuFileController.
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

// UpdateMenuFile implements MenuFileController.
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
