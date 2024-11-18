package req

type CreateRestaurantReq struct {
	Name   string `json:"name" binding:"required,min=3,max=100"`
	UserID uint   `json:"user_id" binding:"required"`
}
