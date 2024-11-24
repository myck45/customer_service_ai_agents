package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/service"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	userService service.UserService
}

// CreateUser implements UserController.
func (u *UserControllerImpl) CreateUser(c *gin.Context) {

	createUserReq := &req.CreateUserReq{}
	if err := c.ShouldBindJSON(createUserReq); err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error binding request")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error binding request: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := u.userService.CreateUser(createUserReq)
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error creating user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error creating user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Msg:    "User created successfully",
		Data:   user.ID,
	}

	c.JSON(http.StatusCreated, res)

}

// DeleteUser implements UserController.
func (u *UserControllerImpl) DeleteUser(c *gin.Context) {

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error parsing id")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error parsing id: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := u.userService.DeleteUser(uint(id)); err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error deleting user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error deleting user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "User deleted successfully",
		Data:   nil,
	}

	c.JSON(http.StatusOK, res)
}

// GetAllUsers implements UserController.
func (u *UserControllerImpl) GetAllUsers(c *gin.Context) {

	users, err := u.userService.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error fetching users")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error fetching users: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "Users fetched successfully",
		Data:   users,
	}

	c.JSON(http.StatusOK, res)
}

// GetUserByEmail implements UserController.
func (u *UserControllerImpl) GetUserByEmail(c *gin.Context) {

	email := c.Param("email")

	user, err := u.userService.GetUserByEmail(email)
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error fetching user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error fetching user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "User fetched successfully",
		Data:   user,
	}

	c.JSON(http.StatusOK, res)
}

// GetUserByID implements UserController.
func (u *UserControllerImpl) GetUserByID(c *gin.Context) {

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error parsing id")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error parsing id: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := u.userService.GetUserByID(uint(id))
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error fetching user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error fetching user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "User fetched successfully",
		Data:   user,
	}

	c.JSON(http.StatusOK, res)
}

// UpdateUser implements UserController.
func (u *UserControllerImpl) UpdateUser(c *gin.Context) {

	updateUserReq := &req.UpdateUserReq{}
	if err := c.ShouldBindJSON(updateUserReq); err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error binding request")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error binding request: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error parsing id")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error parsing id: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := u.userService.UpdateUser(uint(id), updateUserReq); err != nil {
		logrus.WithError(err).Error("[UserControllerImpl] Error updating user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error updating user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "User updated successfully",
		Data:   nil,
	}

	c.JSON(http.StatusOK, res)
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}
