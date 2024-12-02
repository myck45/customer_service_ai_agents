package res

// MenuResponse represents the response to a menu request
// @Description response to a menu request
type MenuResponse struct {
	ID           uint   `json:"id" example:"1" extensions:"x-order=0"`
	RestaurantID uint   `json:"restaurant_id" example:"1" extensions:"x-order=1"`
	ItemName     string `json:"item_name" example:"Completo italiano" extensions:"x-order=2"`
	Description  string `json:"description" example:"Un completo italiano con mayo" extensions:"x-order=3"`
	Price        int    `json:"price" example:"2500" extensions:"x-order=4"`
	Likes        int    `json:"likes" example:"99" extensions:"x-order=5"`
	CreatedAt    string `json:"created_at" example:"2021-07-01T00:00:00Z" extensions:"x-order=6"`
	UpdatedAt    string `json:"updated_at" example:"2021-07-01T00:00:00Z" extensions:"x-order=7"`
}
