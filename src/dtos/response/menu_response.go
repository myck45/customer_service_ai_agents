package response

type MenuResponse struct {
	ID           uint   `json:"id"`
	RestaurantID uint   `json:"restaurant_id"`
	ItemName     string `json:"item_name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
