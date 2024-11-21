package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyectos01-a/bot/dto/req"
	"github.com/proyectos01-a/bot/service"
	"github.com/proyectos01-a/shared/handlers"
	"github.com/sirupsen/logrus"
)

type BotCRUDControllerImpl struct {
	botCRUDService  service.BotCRUDService
	responseHandler handlers.ResponseHandlers
}

// CreateBot implements BotCRUDController.
func (b *BotCRUDControllerImpl) CreateBot(c *gin.Context) {

	createBotRequest := &req.CreateBotReq{}
	if err := c.ShouldBindJSON(createBotRequest); err != nil {
		logrus.WithError(err).Error("*** [CreateBot] Error binding request")
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	bot, err := b.botCRUDService.CreateBot(createBotRequest)
	if err != nil {
		logrus.WithError(err).Error("*** [CreateBot] Error creating bot")
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error creating bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot created successfully", bot.ID)
}

// DeleteBotByID implements BotCRUDController.
func (b *BotCRUDControllerImpl) DeleteBotByID(c *gin.Context) {

	botID := c.Param("id")
	id, err := strconv.ParseUint(botID, 10, 64)
	if err != nil {
		logrus.WithError(err).Error("*** [DeleteBotByID] Error parsing id")
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	if err := b.botCRUDService.DeleteBotByID(uint(id)); err != nil {
		logrus.WithError(err).Error("*** [DeleteBotByID] Error deleting bot")
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error deleting bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot deleted successfully", nil)
}

// GetAllBots implements BotCRUDController.
func (b *BotCRUDControllerImpl) GetAllBots(c *gin.Context) {

	bots, err := b.botCRUDService.GetAllBots()
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching bots", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bots fetched successfully", bots)
}

// GetBotByID implements BotCRUDController.
func (b *BotCRUDControllerImpl) GetBotByID(c *gin.Context) {

	botID := c.Param("id")
	id, err := strconv.ParseUint(botID, 10, 64)
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	bot, err := b.botCRUDService.GetBotByID(uint(id))
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot fetched successfully", bot)
}

// GetBotByRestaurantID implements BotCRUDController.
func (b *BotCRUDControllerImpl) GetBotByRestaurantID(c *gin.Context) {

	restaurantID := c.Query("restaurant_id")
	id, err := strconv.ParseUint(restaurantID, 10, 64)
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	bots, err := b.botCRUDService.GetBotByRestaurantID(uint(id))
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot fetched successfully", bots)
}

// GetBotByWspNumber implements BotCRUDController.
func (b *BotCRUDControllerImpl) GetBotByWspNumber(c *gin.Context) {

	wspNumber := c.Param("whatsapp")

	bot, err := b.botCRUDService.GetBotByWspNumber(wspNumber)
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error fetching bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot fetched successfully", bot)
}

// UpdateBot implements BotCRUDController.
func (b *BotCRUDControllerImpl) UpdateBot(c *gin.Context) {

	botID := c.Param("id")
	id, err := strconv.ParseUint(botID, 10, 64)
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error parsing id", err)
		return
	}

	updateBotRequest := &req.UpdateBotReq{}
	if err := c.ShouldBindJSON(updateBotRequest); err != nil {
		b.responseHandler.HandleError(c, http.StatusBadRequest, "Error binding request", err)
		return
	}

	bot, err := b.botCRUDService.UpdateBot(uint(id), updateBotRequest)
	if err != nil {
		b.responseHandler.HandleError(c, http.StatusInternalServerError, "Error updating bot", err)
		return
	}

	b.responseHandler.HandleSuccess(c, http.StatusOK, "Bot updated successfully", bot)
}

func NewBotCRUDControllerImpl(botCRUDService service.BotCRUDService, responseHandler handlers.ResponseHandlers) BotCRUDController {
	return &BotCRUDControllerImpl{
		botCRUDService:  botCRUDService,
		responseHandler: responseHandler,
	}
}
