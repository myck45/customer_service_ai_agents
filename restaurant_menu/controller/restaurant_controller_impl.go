package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/service"
	_ "github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/handlers"
)

type RestaurantControllerImpl struct {
	restaurantService service.RestaurantService
	responseHandler   handlers.ResponseHandlers
}

// CreateRestaurant godoc
//
//	@Summary		Create a new restaurant
//	@Description	create a new restaurant with the input payload
//	@Tags			restaurant
//	@Accept			json
//	@Produce		json
//	@Param			request				body		req.CreateRestaurantReq		true	"Create Restaurant Request"
//	@Success		200					{object}	dto.BaseResponse{data=uint}	"Restaurant created successfully"
//	@Failure		400					{object}	dto.BaseResponse			"Error binding request"
//	@Failure		500					{object}	dto.BaseResponse			"Error creating restaurant"
//	@Router			/api/v1/restaurant	[post]
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

// DeleteRestaurant godoc
//
//	@Summary		Delete a restaurant
//	@Description	delete a restaurant with the input id
//	@Tags			restaurant
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int					true	"Restaurant ID"
//	@Success		200						{object}	dto.BaseResponse	"Restaurant deleted successfully"
//	@Failure		400						{object}	dto.BaseResponse	"Error parsing id"
//	@Failure		500						{object}	dto.BaseResponse	"Error deleting restaurant"
//	@Router			/api/v1/restaurant/{id}	[delete]
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

// GetAllRestaurants godoc
//
//	@Summary		Get all restaurants
//	@Description	get all restaurants
//	@Tags			restaurant
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	dto.BaseResponse{data=[]res.RestaurantResponse}	"Restaurants fetched successfully"
//	@Failure		500					{object}	dto.BaseResponse								"Error fetching restaurants"
//	@Router			/api/v1/restaurant	[get]
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

// UpdateRestaurant godoc
//
//	@Summary		Update a restaurant
//	@Description	update a restaurant with the input id and payload
//	@Tags			restaurant
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int						true	"Restaurant ID"
//	@Param			request					body		req.UpdateRestaurantReq	true	"Update Restaurant Request"
//	@Success		200						{object}	dto.BaseResponse		"Restaurant updated successfully"
//	@Failure		400						{object}	dto.BaseResponse		"Error binding request"
//	@Failure		500						{object}	dto.BaseResponse		"Error updating restaurant"
//	@Router			/api/v1/restaurant/{id}	[put]
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
