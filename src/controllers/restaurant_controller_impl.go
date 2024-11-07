package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/request"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos/response"
	"github.com/proyectos01-a/RestaurantMenu/src/services"
	"github.com/sirupsen/logrus"
)

type RestaurantControllerImpl struct {
	restaurantService services.RestaurantService
}

// CreateRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) CreateRestaurant(c *gin.Context) {
	createRestaurantReq := &request.CreateRestaurantReq{}
	err := c.ShouldBindJSON(createRestaurantReq)
	if err != nil {
		logrus.WithError(err).Error("Error binding request")
		res := dtos.BaseResponse[string]{
			Code:   400,
			Status: "Bad Request",
			Msg:    fmt.Sprintf("Error binding request: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	restaurant, err := r.restaurantService.CreateRestaurant(createRestaurantReq)
	if err != nil {
		logrus.WithError(err).Error("Error creating restaurant")
		res := dtos.BaseResponse[string]{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    fmt.Sprintf("Error creating restaurant: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	res := dtos.BaseResponse[uint]{
		Code:   200,
		Status: "OK",
		Msg:    "Restaurant created successfully",
		Data:   restaurant.ID,
	}

	c.JSON(200, res)
}

// DeleteRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) DeleteRestaurant(c *gin.Context) {

	restaurantID := c.Param("id")

	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Error parsing id")
		res := dtos.BaseResponse[string]{
			Code:   400,
			Status: "Bad Request",
			Msg:    fmt.Sprintf("Error parsing id: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	err = r.restaurantService.DeleteRestaurant(uint(id))
	if err != nil {
		logrus.WithError(err).Error("Error deleting restaurant")
		res := dtos.BaseResponse[string]{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    fmt.Sprintf("Error deleting restaurant: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	res := dtos.BaseResponse[string]{
		Code:   200,
		Status: "OK",
		Msg:    "Restaurant deleted successfully",
		Data:   "",
	}

	c.JSON(200, res)
}

// GetAllRestaurants implements RestaurantController.
func (r *RestaurantControllerImpl) GetAllRestaurants(c *gin.Context) {
	restaurants, err := r.restaurantService.GetAllRestaurants()
	if err != nil {
		logrus.WithError(err).Error("Error fetching restaurants")
		res := dtos.BaseResponse[string]{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    fmt.Sprintf("Error fetching restaurants: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	res := dtos.BaseResponse[[]response.RestaurantResponse]{
		Code:   200,
		Status: "OK",
		Msg:    "Restaurants fetched successfully",
		Data:   restaurants.Restaurants,
	}

	c.JSON(200, res)

}

// GetRestaurantByID implements RestaurantController.
func (r *RestaurantControllerImpl) GetRestaurantByID(c *gin.Context) {
	restaurantID := c.Param("id")

	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Error parsing id")
		res := dtos.BaseResponse[string]{
			Code:   400,
			Status: "Bad Request",
			Msg:    fmt.Sprintf("Error parsing id: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	restaurant, err := r.restaurantService.GetRestaurantByID(uint(id))
	if err != nil {
		logrus.WithError(err).Error("Error fetching restaurant")
		res := dtos.BaseResponse[string]{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    fmt.Sprintf("Error fetching restaurant: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	res := dtos.BaseResponse[*response.RestaurantResponse]{
		Code:   200,
		Status: "OK",
		Msg:    "Restaurant fetched successfully",
		Data:   restaurant,
	}

	c.JSON(200, res)
}

// UpdateRestaurant implements RestaurantController.
func (r *RestaurantControllerImpl) UpdateRestaurant(c *gin.Context) {
	restaurantID := c.Param("id")

	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("Error parsing id")
		res := dtos.BaseResponse[string]{
			Code:   400,
			Status: "Bad Request",
			Msg:    fmt.Sprintf("Error parsing id: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	updateRestaurantReq := &request.UpdateRestaurantReq{}

	err = c.ShouldBindJSON(updateRestaurantReq)
	if err != nil {
		logrus.WithError(err).Error("Error binding request")
		res := dtos.BaseResponse[string]{
			Code:   400,
			Status: "Bad Request",
			Msg:    fmt.Sprintf("Error binding request: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	restaurant, err := r.restaurantService.UpdateRestaurant(uint(id), updateRestaurantReq)
	if err != nil {
		logrus.WithError(err).Error("Error updating restaurant")
		res := dtos.BaseResponse[string]{
			Code:   500,
			Status: "Internal Server Error",
			Msg:    fmt.Sprintf("Error updating restaurant: %s", err.Error()),
			Data:   "",
		}

		c.JSON(200, res)
		return
	}

	res := dtos.BaseResponse[*response.RestaurantResponse]{
		Code:   200,
		Status: "OK",
		Msg:    "Restaurant updated successfully",
		Data:   restaurant,
	}

	c.JSON(200, res)
}

func NewRestaurantControllerImpl(restaurantService services.RestaurantService) RestaurantController {
	return &RestaurantControllerImpl{
		restaurantService: restaurantService,
	}
}
