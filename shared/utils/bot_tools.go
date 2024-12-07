package utils

import "github.com/sashabaranov/go-openai"

type BotTools interface {
	// OCR for menu items from image
	GetMenuItemsFromImage() *openai.FunctionDefinition

	// Capture user order
	GetUserOrder() *openai.FunctionDefinition

	// Delete user order
	DeleteUserOrder() *openai.FunctionDefinition
}
