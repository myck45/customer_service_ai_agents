package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/service"
)

type UserControllerImpl struct {
	userService     service.UserService
	responseHandler handlers.ResponseHandlers
}

// CreateUser implements UserController.
func (u *UserControllerImpl) CreateUser(c *gin.Context) {

	createUserReq := &req.CreateUserReq{}
	if err := c.ShouldBindJSON(createUserReq); err != nil {
		u.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	user, err := u.userService.CreateUser(createUserReq)
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error creating user", err)

	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "User created successfully", user.ID)

}

// DeleteUser implements UserController.
func (u *UserControllerImpl) DeleteUser(c *gin.Context) {

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := u.userService.DeleteUser(uint(id)); err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error deleting user", err)
		return
	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "User deleted successfully", nil)
}

// GetAllUsers implements UserController.
func (u *UserControllerImpl) GetAllUsers(c *gin.Context) {

	users, err := u.userService.GetAllUsers()
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching users", err)
		return
	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "Users fetched successfully", users)
}

// GetUserByEmail implements UserController.
func (u *UserControllerImpl) GetUserByEmail(c *gin.Context) {

	email := c.Param("email")

	user, err := u.userService.GetUserByEmail(email)
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching user", err)
		return
	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "User fetched successfully", user)
}

// GetUserByID implements UserController.
func (u *UserControllerImpl) GetUserByID(c *gin.Context) {

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	user, err := u.userService.GetUserByID(uint(id))
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching user", err)
		return
	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "User fetched successfully", user)
}

// UpdateUser implements UserController.
func (u *UserControllerImpl) UpdateUser(c *gin.Context) {

	updateUserReq := &req.UpdateUserReq{}
	if err := c.ShouldBindJSON(updateUserReq); err != nil {
		u.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		u.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := u.userService.UpdateUser(uint(id), updateUserReq); err != nil {
		u.responseHandler.HandleError(c, http.StatusInternalServerError, "Error updating user", err)
		return
	}

	u.responseHandler.HandleSuccess(c, http.StatusOK, "User updated successfully", nil)
}

func NewUserControllerImpl(userService service.UserService, responseHandler handlers.ResponseHandlers) UserController {
	return &UserControllerImpl{
		userService:     userService,
		responseHandler: responseHandler,
	}
}
