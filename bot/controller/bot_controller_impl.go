package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/bot/service"
	"github.com/proyectos01-a/shared/handlers"
)

type BotControllerImpl struct {
	botService      service.BotService
	responseHandler handlers.ResponseHandlers
}

// BotResponse implements BotController.
func (b *BotControllerImpl) BotResponse(c *gin.Context) {

	// Capture twilio request
	userWspNumber := c.PostForm("From")
	botWspNumber := c.PostForm("To")
	userMessage := c.PostForm("Body")

	// Build twilio request
	twilioReq := &req.TwilioWebhook{
		To:   botWspNumber,
		From: userWspNumber,
		Body: userMessage,
	}

	if err := c.ShouldBind(twilioReq); err != nil {
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	if err := b.botService.BotResponse(twilioReq); err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error processing bot response", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot response processed successfully", nil)
}

func NewBotControllerImpl(botService service.BotService, responseHandler handlers.ResponseHandlers) BotController {
	return &BotControllerImpl{
		botService:      botService,
		responseHandler: responseHandler,
	}
}
