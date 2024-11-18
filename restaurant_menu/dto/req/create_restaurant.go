package req

type CreateRestaurantReq struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
