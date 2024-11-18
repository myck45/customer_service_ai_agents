package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/service"
	"github.com/proyectos01-a/shared/handlers"
)

type RestaurantControllerImpl struct {
	restaurantService service.RestaurantService
	responseHandler   handlers.ResponseHandlers
}

// CreateRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) CreateRestaurant(c *gin.Context) {
	createRestaurantReq := &req.CreateRestaurantReq{}
	if err := c.ShouldBindJSON(createRestaurantReq); err != nil {
		r.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	restaurant, err := r.restaurantService.CreateRestaurant(createRestaurantReq)
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusInternalServerError, "Error creating restaurant", err)
	}

	r.responseHandler.HandleSuccess(c, http.StatusOK, "Restaurant created successfully", restaurant.ID)
}

// DeleteRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) DeleteRestaurant(c *gin.Context) {

	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := r.restaurantService.DeleteRestaurant(uint(id)); err != nil {
		r.responseHandler.HandleError(c, http.StatusInternalServerError, "Error deleting restaurant", err)
		return
	}

	r.responseHandler.HandleSuccess(c, http.StatusOK, "Restaurant deleted successfully", nil)
}

// GetAllRestaurants implements RestaurantController.
func (r *RestaurantControllerImpl) GetAllRestaurants(c *gin.Context) {
	restaurants, err := r.restaurantService.GetAllRestaurants()
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching restaurants", err)
		return
	}

	r.responseHandler.HandleSuccess(c, http.StatusOK, "Restaurants fetched successfully", restaurants)
}

// GetRestaurantByID implements RestaurantController.
func (r *RestaurantControllerImpl) GetRestaurantByID(c *gin.Context) {

	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	restaurant, err := r.restaurantService.GetRestaurantByID(uint(id))
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching restaurant", err)
		return
	}

	r.responseHandler.HandleSuccess(c, http.StatusOK, "Restaurant fetched successfully", restaurant)
}

// UpdateRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) UpdateRestaurant(c *gin.Context) {

	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	updateRestaurantReq := &req.UpdateRestaurantReq{}
	if err := c.ShouldBindJSON(updateRestaurantReq); err != nil {
		r.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	restaurant, err := r.restaurantService.UpdateRestaurant(uint(id), updateRestaurantReq)
	if err != nil {
		r.responseHandler.HandleError(c, http.StatusInternalServerError, "Error updating restaurant", err)
		return
	}

	r.responseHandler.HandleSuccess(c, http.StatusOK, "Restaurant updated successfully", restaurant)
}

func NewRestaurantControllerImpl(restaurantService service.RestaurantService, responseHandler handlers.ResponseHandlers) RestaurantController {
	return &RestaurantControllerImpl{
		restaurantService: restaurantService,
		responseHandler:   responseHandler,
	}
}
