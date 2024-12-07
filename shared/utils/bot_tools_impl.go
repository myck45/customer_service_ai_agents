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
	// schema, err := jsonschema.GenerateSchemaForType(schemas.UserOrderFunctionSchema{})
	// if err != nil {
	// 	logrus.WithError(err).Error("[getUserOrder] failed to generate schema")
	// 	return nil
	// }

	schema := &jsonschema.Definition{
		Type:     jsonschema.Object,
		Required: []string{"menu_items", "delivery_address", "user_name", "phone_number", "payment_method"},
		Properties: map[string]jsonschema.Definition{
			"menu_items": {
				Type: jsonschema.Array,
				Items: &jsonschema.Definition{
					Type: jsonschema.Object,
					Properties: map[string]jsonschema.Definition{
						"item_name": {
							Type:        jsonschema.String,
							Description: "Nombre del ítem del menú",
						},
						"quantity": {
							Type:        jsonschema.Integer,
							Description: "Cantidad del ítem del menú solicitada por el usuario",
						},
						"price": {
							Type:        jsonschema.Number,
							Description: "Precio del ítem del menú",
						},
					},
					Required:             []string{"item_name", "quantity", "price"},
					AdditionalProperties: false,
				},
				Description: "Lista de ítems del menú solicitados por el usuario",
			},
			"delivery_address": {
				Type:        jsonschema.String,
				Description: "Dirección de entrega del pedido",
			},
			"user_name": {
				Type:        jsonschema.String,
				Description: "Nombre del usuario que realiza el pedido",
			},
			"phone_number": {
				Type:        jsonschema.String,
				Description: "Número de teléfono del usuario que realiza el pedido",
			},
			"payment_method": {
				Type:        jsonschema.String,
				Description: "Método de pago del pedido",
				Enum:        []string{"efectivo", "transferencia"},
			},
		},
		AdditionalProperties: false,
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

// DeleteUserOrder implements BotTools.
func (b *BotToolsImpl) DeleteUserOrder() *openai.FunctionDefinition {
	// schema, err := jsonschema.GenerateSchemaForType(schemas.DeleteUserOrderFunctionSchema{})
	// if err != nil {
	// 	logrus.WithError(err).Error("[deleteUserOrder] failed to generate schema")
	// 	return nil
	// }

	schema := &jsonschema.Definition{
		Type:     jsonschema.Object,
		Required: []string{"order_code", "user_confirm"},
		Properties: map[string]jsonschema.Definition{
			"order_code": {
				Type:        jsonschema.String,
				Description: "Código del pedido a eliminar",
			},
			"user_confirmation": {
				Type:        jsonschema.String,
				Description: "Confirmación del usuario para eliminar el pedido",
				Enum:        []string{"si", "no"},
			},
		},
		AdditionalProperties: false,
	}

	return &openai.FunctionDefinition{
		Name:        "delete_user_order",
		Description: "Elimina el pedido del usuario, es necesario que el usuario proporcione el código del pedido a eliminar. Debes preguntar al usuario si está seguro de eliminar el pedido con el código proporcionado.",
		Parameters:  schema,
		Strict:      true,
	}
}

func NewBotTools() BotTools {
	return &BotToolsImpl{}
}
