package request

type UpdateBotReq struct {
	Name      string `json:"name" binding:"required,min=1,max=100"`
	Identity  string `json:"identity" binding:"required,min=1,max=100"`
	WspNumber string `json:"wsp_number" binding:"required,min=1,max=100"`
}
