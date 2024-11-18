package res

type BotResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Identity     string `json:"identity"`
	WspNumber    string `json:"wsp_number"`
	RestaurantID uint   `json:"restaurant_id"`
}
