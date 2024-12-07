package req

type DeleteUserOrderRequest struct {
	OrderCode        string `json:"order_code"`
	UserConfirmation string `json:"user_confirmation"`
}
