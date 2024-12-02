package req

// CreateMenuReq represents the request to create a new menu item
// @Description create a new entry in the menu table
type CreateMenuReq struct {
	RestaurantID uint   `json:"restaurant_id" binding:"required" example:"1" extensions:"x-order=0"`                                         // RestaurantID is the ID of the restaurant that the menu item belongs to
	ItemName     string `json:"item_name" binding:"required,min=1,max=100" example:"Completo italiano" extensions:"x-order=1"`               // ItemName is the name of the menu item
	Description  string `json:"description" binding:"required,min=1,max=255" example:"Un completo italiano con mayo" extensions:"x-order=2"` // Description is the description of the menu item
	Likes        int    `json:"likes" binding:"required,min=0" example:"99" extensions:"x-order=3"`                                          // Likes is the number of likes of the menu item
	Price        int    `json:"price" binding:"required,min=0" example:"2500" extensions:"x-order=4"`                                        // Price is the price of the menu item
}
