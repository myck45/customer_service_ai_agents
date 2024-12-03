package utils

import openai "github.com/sashabaranov/go-openai"

type BotToolsImpl struct{}

// getUserOrder implements BotTools.
func (b *BotToolsImpl) getUserOrder() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        "get_user_order",
		Description: "Obtiene el pedido del usuario, es necesario que el usuario proporcione los ítems del menú solicitados, la dirección de entrega, su nombre, número de teléfono y método de pago.",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"menu_items": map[string]interface{}{
					"type":        "array",
					"description": "Lista de ítems del menú solicitados por el usuario.",
					"items": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"item_name": map[string]interface{}{
								"type":        "string",
								"description": "Nombre del ítem del menú.",
							},
							"quantity": map[string]interface{}{
								"type":        "number",
								"description": "Cantidad del ítem del menú solicitada por el usuario.",
							},
						},
						"required": []string{"item_name", "quantity"},
					},
				},
				"delivery_address": map[string]interface{}{
					"type":        "string",
					"description": "Dirección de entrega del pedido.",
				},
				"user_name": map[string]interface{}{
					"type":        "string",
					"description": "Nombre del usuario que realiza el pedido.",
				},
				"phone_number": map[string]interface{}{
					"type":        "string",
					"description": "Número de teléfono del usuario que realiza el pedido.",
				},
				"payment_method": map[string]interface{}{
					"type":        "string",
					"description": "Método de pago del pedido.",
				},
			},
			"required": []string{"menu_items", "delivery_address", "user_name", "phone_number", "payment_method"},
		},
	}
}

// getMenuItemsFromImage implements BotTools.
func (b *BotToolsImpl) getMenuItemsFromImage() openai.FunctionDefinition {
	return openai.FunctionDefinition{
		Name:        "extract_menu_items",
		Description: "Extrae los ítems del menú de una imagen de un menú.",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"items": map[string]interface{}{
					"type":        "array",
					"description": "Lista de ítems del menú extraídos de la imagen.",
					"items": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"item_name": map[string]interface{}{
								"type":        "string",
								"description": "Nombre del ítem del menú.",
							}, // item_name - type string
							"description": map[string]interface{}{
								"type":        "string",
								"description": "Descripción del ítem del menú.",
							}, // description - type string
							"price": map[string]interface{}{
								"type":        "number",
								"description": "Precio del ítem del menú.",
							}, // price - type number
						},
					}, // item del menu - type object
					"required": []string{"item_name", "description", "price"},
				}, // array de items del menu - type array
			},
			"required": []string{"items"},
		}, // parametros - type object
	}
}

func NewBotTools() BotTools {
	return &BotToolsImpl{}
}
