package res

// RestaurantResponse represents the response to a restaurant request
// @Description response to a restaurant request
type RestaurantResponse struct {
	ID        uint   `json:"id" example:"1" extensions:"x-order=0"`
	Name      string `json:"name" example:"La Piccola Italia" extensions:"x-order=1"`
	CreatedAt string `json:"created_at" example:"2021-07-01T00:00:00Z" extensions:"x-order=2"`
	UpdatedAt string `json:"updated_at" example:"2021-07-01T00:00:00Z" extensions:"x-order=3"`
}
