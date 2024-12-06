package schemas

type MenuItemSchema struct {
	ItemName string `json:"item_name" description:"Nombre del ítem del menú" required:"true"`
	Quantity int    `json:"quantity" description:"Cantidad del ítem del menú solicitada por el usuario" required:"true"`
	Price    int    `json:"price" description:"Precio del ítem del menú" required:"true"`
}

type UserOrderFunctionSchema struct {
	MenuItems       []MenuItemSchema `json:"menu_items" description:"Lista de ítems del menú solicitados por el usuario" required:"true"`
	DeliveryAddress string           `json:"delivery_address" description:"Dirección de entrega del pedido" required:"true"`
	UserName        string           `json:"user_name" description:"Nombre del usuario que realiza el pedido" required:"true"`
	PhoneNumber     string           `json:"phone_number" description:"Número de teléfono del usuario que realiza el pedido" required:"true"`
	PaymentMethod   string           `json:"payment_method" description:"Método de pago del pedido" required:"true" enum:"efectivo,transferencia"`
}
