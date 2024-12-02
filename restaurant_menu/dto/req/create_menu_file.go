package req

// CreateMenuFileReq represents the request to create a new menu file
// @Description create a new entry in the menu files table
type CreateMenuFileReq struct {
	RestaurantID uint   `json:"restaurant_id" binding:"required" example:"1" extensions:"x-order=0"`    // RestaurantID is the ID of the restaurant that the menu file belongs to
	FileName     string `json:"file_name" binding:"required" example:"menu.pdf" extensions:"x-order=1"` // FileName is the name of the file
	FileType     string `json:"file_type" binding:"required" example:"pdf" extensions:"x-order=2"`      // FileType is the type of the file
}
