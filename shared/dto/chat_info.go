package dto

type ChatInfoRequest struct {
	BotWspNumber    string `json:"bot_wsp_number"`
	SenderWspNumber string `json:"sender_wsp_number"`
	RestaurantID    uint   `json:"restaurant_id"`
}
