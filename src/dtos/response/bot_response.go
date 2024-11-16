package response

type BotResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	WspNumber    string `json:"wsp_number"`
	RestaurantID uint   `json:"restaurant_id"`
}
