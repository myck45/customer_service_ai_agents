package dto

// BaseResponse struct
// @Description base response struct for JSON responses
type BaseResponse struct {
	Code   int         `json:"code" example:"200" extensions:"x-order=0"`
	Status string      `json:"status" example:"success" extensions:"x-order=1"`
	Msg    string      `json:"msg" example:"success" extensions:"x-order=2"`
	Data   interface{} `json:"data" example:"" extensions:"x-order=3"`
}
