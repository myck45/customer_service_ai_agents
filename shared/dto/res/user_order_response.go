package res

type UserOrderResponse struct {
	DeliveryAddress string          `json:"delivery_address"`
	UserName        string          `json:"user_name"`
	PhoneNumber     string          `json:"phone_number"`
	PaymentMethod   string          `json:"payment_method"`
	MenuItems       []UserOrderItem `json:"menu_items"`
}

type UserOrderItem struct {
	ItemName  string `json:"item_name"`
	Quantity  int    `json:"quantity"`
	ItemPrice int    `json:"item_price"`
}
