package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/service"
	"github.com/sirupsen/logrus"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

// Login implements AuthController.
func (a *AuthControllerImpl) Login(c *gin.Context) {

	loginReq := &req.LoginRequest{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		logrus.WithError(err).Error("[AuthControllerImpl] Error binding request")
		res := dto.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Msg:    fmt.Sprintf("Error binding request: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := a.authService.UserLogin(loginReq)
	if err != nil {
		logrus.WithError(err).Error("[AuthControllerImpl] Error logging in user")
		res := dto.BaseResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Msg:    fmt.Sprintf("Error logging in user: %v", err),
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := dto.BaseResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Msg:    "User logged in successfully",
		Data:   token,
	}

	c.JSON(http.StatusOK, res)
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}
