package utils

import "github.com/sashabaranov/go-openai"

type BotTools interface {
	getMenuItemsFromImage() openai.FunctionDefinition
	getUserOrder() openai.FunctionDefinition
}
