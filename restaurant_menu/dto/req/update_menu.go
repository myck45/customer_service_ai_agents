package req

// UpdateMenuReq represents the request to update a menu item
// @Description update an entry in the menu table
type UpdateMenuReq struct {
	ItemName    string `json:"item_name" binding:"required,min=1,max=100" example:"Completo italiano" extensions:"x-order=0"`
	Description string `json:"description" binding:"required,min=1,max=255" example:"Un completo italiano con mayo" extensions:"x-order=1"`
	Price       int    `json:"price" binding:"required,min=0" example:"2500" extensions:"x-order=2"`
	Likes       int    `json:"likes" binding:"required,min=0" example:"99" extensions:"x-order=3"`
}
