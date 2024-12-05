package utils

import "github.com/sashabaranov/go-openai"

type BotTools interface {
	GetMenuItemsFromImage() *openai.FunctionDefinition
	GetUserOrder() *openai.FunctionDefinition
}
