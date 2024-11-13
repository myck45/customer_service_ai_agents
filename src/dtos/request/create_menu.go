package request

type CreateMenuReq struct {
	RestaurantID uint   `json:"restaurant_id" binding:"required"`
	ItemName     string `json:"item_name" binding:"required,min=1,max=100"`
	Description  string `json:"description" binding:"required,min=1,max=255"`
	Likes        int    `json:"likes" binding:"required,min=0"`
	Price        int    `json:"price" binding:"required,min=0"`
}
