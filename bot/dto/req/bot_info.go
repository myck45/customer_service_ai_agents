package req

type BotInfo struct {
	BotName      string `json:"bot_name"`
	BotIdentity  string `json:"bot_identity"`
	RestaurantID uint   `json:"restaurant_id"`
}
