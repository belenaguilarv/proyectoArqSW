package dto

type OrderDto struct {
	OrderId    int        `json:"order_id"`
	Date       string     `json:"date"`
	TotalPrice int        `json:"total_price"`
	UserId     int        `json:"user_id"`
	Detail     DetailsDto `json:"details"` //no se si es necesario
}

type OrdersDto []OrderDto
