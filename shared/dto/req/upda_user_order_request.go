package req

type UpdateUserOrderRequest struct {
	OrderCode        string            `json:"order_code"`
	UserConfirmation string            `json:"user_confirmation"`
	MenuItems        []MenuItemRequest `json:"menu_items"`
	DeliveryAddress  string            `json:"delivery_address"`
	UserName         string            `json:"user_name"`
	PhoneNumber      string            `json:"phone_number"`
	PaymentMethod    string            `json:"payment_method"`
}
