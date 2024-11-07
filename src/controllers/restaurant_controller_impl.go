package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/services"
	"github.com/sirupsen/logrus"
)

type RestaurantControllerImpl struct {
	restaurantService services.RestaurantService
}

func NewRestaurantControllerImpl(restaurantService services.RestaurantService) RestaurantController {
	return &RestaurantControllerImpl{
		restaurantService: restaurantService,
	}
}

func (r *RestaurantControllerImpl) CreateRestaurant(c *gin.Context) {
	createRestaurantReq := &request.CreateRestaurantReq{}
	if err := c.ShouldBindJSON(createRestaurantReq); err != nil {
		r.handleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	restaurant, err := r.restaurantService.CreateRestaurant(createRestaurantReq)
	if err != nil {
		r.handleError(c, http.StatusInternalServerError, "Error creating restaurant", err)
		return
	}

	r.handleSuccess(c, http.StatusOK, "Restaurant created successfully", restaurant.ID)
}

func (r *RestaurantControllerImpl) DeleteRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.handleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := r.restaurantService.DeleteRestaurant(uint(id)); err != nil {
		r.handleError(c, http.StatusInternalServerError, "Error deleting restaurant", err)
		return
	}

	r.handleSuccess(c, http.StatusOK, "Restaurant deleted successfully", nil)
}

func (r *RestaurantControllerImpl) GetAllRestaurants(c *gin.Context) {
	restaurants, err := r.restaurantService.GetAllRestaurants()
	if err != nil {
		r.handleError(c, http.StatusInternalServerError, "Error fetching restaurants", err)
		return
	}

	r.handleSuccess(c, http.StatusOK, "Restaurants fetched successfully", restaurants.Restaurants)
}

func (r *RestaurantControllerImpl) GetRestaurantByID(c *gin.Context) {
	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.handleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	restaurant, err := r.restaurantService.GetRestaurantByID(uint(id))
	if err != nil {
		r.handleError(c, http.StatusInternalServerError, "Error fetching restaurant", err)
		return
	}

	r.handleSuccess(c, http.StatusOK, "Restaurant fetched successfully", restaurant)
}

func (r *RestaurantControllerImpl) UpdateRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		r.handleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	updateRestaurantReq := &request.UpdateRestaurantReq{}
	if err := c.ShouldBindJSON(updateRestaurantReq); err != nil {
		r.handleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	restaurant, err := r.restaurantService.UpdateRestaurant(uint(id), updateRestaurantReq)
	if err != nil {
		r.handleError(c, http.StatusInternalServerError, "Error updating restaurant", err)
		return
	}

	r.handleSuccess(c, http.StatusOK, "Restaurant updated successfully", restaurant)
}

func (r *RestaurantControllerImpl) handleError(c *gin.Context, statusCode int, message string, err error) {
	logrus.WithError(err).Error(message)
	res := dtos.BaseResponse[string]{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    fmt.Sprintf("%s: %s", message, err.Error()),
		Data:   "",
	}
	c.JSON(statusCode, res)
}

func (r *RestaurantControllerImpl) handleSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	res := dtos.BaseResponse[interface{}]{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    message,
		Data:   data,
	}
	c.JSON(statusCode, res)
}
