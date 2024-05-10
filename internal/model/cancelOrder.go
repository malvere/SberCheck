package model

// url: https://megamarket.ru/api/mobile/v1/orderService/order/cancel
type CancelOrderRequest struct {
	OrderID               string   `json:"orderId"`
	CancellationReasonIds []string `json:"cancellationReasonIds"` //r-0003
	Auth                  AuthBody `json:"auth"`
}

type CancelOrderResponse struct {
	Success string `json:"success"`
	Error   any    `json:"error"`
}
