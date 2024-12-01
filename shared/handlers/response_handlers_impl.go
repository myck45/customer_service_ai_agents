package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/shared/dto"
	"github.com/sirupsen/logrus"
)

type ResponseHandlersImpl struct{}

// HandleError implements ResponseHandlers.
func (r *ResponseHandlersImpl) HandleError(c *gin.Context, statusCode int, message string, err error) {
	logrus.WithError(err).Error(message)
	res := dto.BaseResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    fmt.Sprintf("%s: %s", message, err.Error()),
		Data:   "",
	}
	c.JSON(statusCode, res)
}

// HandleSuccess implements ResponseHandlers.
func (r *ResponseHandlersImpl) HandleSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	res := dto.BaseResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Msg:    message,
		Data:   data,
	}
	c.JSON(statusCode, res)
}

func NewResponseHandlersImpl() ResponseHandlers {
	return &ResponseHandlersImpl{}
}
