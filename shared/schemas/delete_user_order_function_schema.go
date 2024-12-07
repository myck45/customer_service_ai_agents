package schemas

type DeleteUserOrderFunctionSchema struct {
	OrderCode string `json:"order_code" description:"Código del pedido a eliminar" required:"true"`
}
