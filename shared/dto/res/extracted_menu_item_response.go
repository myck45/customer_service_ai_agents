package res

type ExtractedMenuItemResponse struct {
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
