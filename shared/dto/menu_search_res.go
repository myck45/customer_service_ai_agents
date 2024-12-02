package dto

// MenuSearchResponse struct
// @Description struct for menu semantic search response
type MenuSearchResponse struct {
	ID          uint    `json:"id" example:"1" extensions:"x-order=0"`
	ItemName    string  `json:"item_name" example:"Nasi Goreng" extensions:"x-order=1"`
	Price       int     `json:"price" example:"15000" extensions:"x-order=2"`
	Description string  `json:"description" example:"Nasi goreng spesial" extensions:"x-order=3"`
	Likes       int     `json:"likes" example:"100" extensions:"x-order=4"`
	Similarity  float32 `json:"similarity" example:"0.8" extensions:"x-order=5"`
}
