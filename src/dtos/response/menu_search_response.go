package response

type MenuSearchResponse struct {
	ID          uint    `json:"id"`
	ItemName    string  `json:"item_name"`
	Price       int     `json:"price"`
	Description string  `json:"description"`
	Similarity  float32 `json:"similarity"`
}
