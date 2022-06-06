package dto

type OrderDto struct {
	OrderId    int    `json:"order_id"`
	Date       string `json:"date"`
	TotalPrice int    `json:"total_price"`
	UserId     int    `json:"user_id"`
	//Details     DetailsDto `json:"details"` // no creo que sea necesario
}

type OrdersDto []OrderDto
