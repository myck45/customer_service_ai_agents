package controllers

import "github.com/gin-gonic/gin"

type ResponseHandlers interface {
	HandleError(c *gin.Context, statusCode int, message string, err error)
	HandleSuccess(c *gin.Context, statusCode int, message string, data interface{})
}
