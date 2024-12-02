package res

// MenuFileResponse represents the response to a menu file request
// @Description response to a menu file request
type MenuFileResponse struct {
	ID           uint   `json:"id" example:"1" extensions:"x-order=0"`
	FileName     string `json:"file_name" example:"menu.pdf" extensions:"x-order=1"`
	FilePath     string `json:"file_path" example:"/path/to/menu.pdf" extensions:"x-order=2"`
	FileType     string `json:"file_type" example:"application/pdf" extensions:"x-order=3"`
	FileSize     int64  `json:"file_size" example:"1024" extensions:"x-order=4"`
	RestaurantID uint   `json:"restaurant_id" example:"1" extensions:"x-order=5"`
}
