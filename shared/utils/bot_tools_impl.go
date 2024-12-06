package utils

import (
	"github.com/proyectos01-a/shared/schemas"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/sirupsen/logrus"
)

type BotToolsImpl struct{}

// getUserOrder implements BotTools.
func (b *BotToolsImpl) GetUserOrder() *openai.FunctionDefinition {
	schema, err := jsonschema.GenerateSchemaForType(schemas.UserOrderFunctionSchema{})
	if err != nil {
		logrus.WithError(err).Error("[getUserOrder] failed to generate schema")
		return nil
	}

	return &openai.FunctionDefinition{
		Name:        "get_user_order",
		Description: "Obtiene el pedido del usuario, es necesario que el usuario proporcione los ítems del menú solicitados, la dirección de entrega, su nombre, número de teléfono y método de pago, el método de pago solo puede ser efectivo o transferencia.",
		Parameters:  schema,
		Strict:      true,
	}
}

// getMenuItemsFromImage implements BotTools.
func (b *BotToolsImpl) GetMenuItemsFromImage() *openai.FunctionDefinition {
	schema, err := jsonschema.GenerateSchemaForType(schemas.MenuItemsFromImageFunctionSchema{})
	if err != nil {
		logrus.WithError(err).Error("[getMenuItemsFromImage] failed to generate schema")
		return nil
	}

	return &openai.FunctionDefinition{
		Name:        "get_menu_items_from_image",
		Description: "Extrae los ítems del menú de una imagen de un menú.",
		Parameters:  schema,
		Strict:      true,
	}
}

func NewBotTools() BotTools {
	return &BotToolsImpl{}
}
