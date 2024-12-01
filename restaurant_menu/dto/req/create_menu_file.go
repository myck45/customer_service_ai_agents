package req

type CreateMenuFileReq struct {
	RestaurantID uint   `json:"restaurant_id" binding:"required"`
	FileName     string `json:"file_name" binding:"required"`
	FileType     string `json:"file_type" binding:"required"`
}
