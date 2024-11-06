package request

type UpdateMenuReq struct {
	ItemName    string `json:"item_name" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"required,min=1,max=255"`
	Price       int    `json:"price" binding:"required,min=0"`
}
