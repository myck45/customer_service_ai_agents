package request

type UpdateRestaurantReq struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
