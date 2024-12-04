package schemas

type MenuItemFromImageSchema struct {
	ItemName    string  `json:"item_name" description:"Nombre del ítem del menú" required:"true"`
	Description string  `json:"description" description:"Descripción del ítem del menú" required:"true"`
	Price       float64 `json:"price" description:"Precio del ítem del menú" required:"true"`
}

type MenuItemsFromImageFunctionSchema struct {
	Items []MenuItemFromImageSchema `json:"items" description:"Lista de ítems del menú extraídos de la imagen" required:"true"`
}
