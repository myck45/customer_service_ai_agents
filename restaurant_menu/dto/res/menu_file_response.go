package res

type MenuFileResponse struct {
	ID           uint   `json:"id"`
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	FileType     string `json:"file_type"`
	FileSize     int64  `json:"file_size"`
	RestaurantID uint   `json:"restaurant_id"`
}
