package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/service"
)

type AuthControllerImpl struct {
	authService     service.AuthService
	responseHandler handlers.ResponseHandlers
}

// Login implements AuthController.
func (a *AuthControllerImpl) Login(c *gin.Context) {

	loginReq := &req.LoginRequest{}
	if err := c.ShouldBindJSON(loginReq); err != nil {
		a.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	token, err := a.authService.UserLogin(loginReq)
	if err != nil {
		a.responseHandler.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	a.responseHandler.HandleSuccess(c, http.StatusOK, "User logged in successfully", token)
}

func NewAuthControllerImpl(authService service.AuthService, responseHandler handlers.ResponseHandlers) AuthController {
	return &AuthControllerImpl{
		authService:     authService,
		responseHandler: responseHandler,
	}
}
