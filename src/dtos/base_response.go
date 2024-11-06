package dtos

type BaseResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
}
