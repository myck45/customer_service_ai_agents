package handlers

import "github.com/proyectos01-a/shared/dto"

type BotToolsHandler interface {
	// BotToolsHandler handle the user order tool call from bot
	HandleGetUserOrder(data string, chatInfo dto.ChatInfoRequest) (string, error)

	// BotToolsHandler handle the menu items from image tool call from bot
	HandleGetMenuItemsFromImage(data string, restaurantID uint) error
}
