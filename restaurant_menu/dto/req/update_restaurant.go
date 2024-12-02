package req

// UpdateRestaurantReq represents the request to update a restaurant
// @Description update an entry in the restaurant table
type UpdateRestaurantReq struct {
	Name string `json:"name" binding:"required,min=3,max=100" example:"La Piccola Italia" extensions:"x-order=0"` // Name is the name of the restaurant
}
