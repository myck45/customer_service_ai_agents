package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/RestaurantMenu/src/dtos"
	"github.com/sirupsen/logrus"
)

type ResponseHandlersImpl struct{}

// HandleError implements ResponseHandlers.
func (r *ResponseHandlersImpl) HandleError(c *gin.Context, statusCode int, message string, err error) {
	logrus.WithError(err).Error(message)
	res := dtos.BaseResponse[string]{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    fmt.Sprintf("%s: %s", message, err.Error()),
		Data:   "",
	}
	c.JSON(http.StatusOK, res)
}

// HandleSuccess implements ResponseHandlers.
func (r *ResponseHandlersImpl) HandleSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	res := dtos.BaseResponse[interface{}]{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    message,
		Data:   data,
	}
	c.JSON(http.StatusOK, res)
}

func NewResponseHandlersImpl() ResponseHandlers {
	return &ResponseHandlersImpl{}
}
