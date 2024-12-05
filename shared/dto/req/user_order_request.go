package req

type MenuItemRequest struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type UserOrderRequest struct {
	MenuItems       []MenuItemRequest `json:"menu_items"`
	DeliveryAddress string            `json:"delivery_address"`
	UserName        string            `json:"user_name"`
	PhoneNumber     string            `json:"phone_number"`
	PaymentMethod   string            `json:"payment_method"`
}
