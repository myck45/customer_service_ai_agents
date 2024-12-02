package req

// CreateRestaurantReq represents the request to create a new restaurant
// @Description create a new entry in the restaurant table
type CreateRestaurantReq struct {
	Name   string `json:"name" binding:"required,min=3,max=100" example:"La Piccola Italia" extensions:"x-order=0"` // Name is the name of the restaurant
	UserID uint   `json:"user_id" binding:"required" example:"1" extensions:"x-order=1"`                            // UserID is the ID of the user that owns the restaurant
}
