package handlers

import (
	"github.com/proyectos01-a/shared/dto"
	"github.com/proyectos01-a/shared/models"
)

type BotToolsHandler interface {
	// BotToolsHandler handle the user order tool call from bot
	HandleGetUserOrder(data string, chatInfo dto.ChatInfoRequest) (*models.UserOrder, error)

	// BotToolsHandler handle the delete user order tool call from bot
	HandleDeleteUserOrder(data string, chatInfo dto.ChatInfoRequest) (string, error)

	// BotToolsHandler handle the menu items from image tool call from bot
	HandleGetMenuItemsFromImage(data string, restaurantID uint) error
}
