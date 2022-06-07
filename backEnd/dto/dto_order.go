package dto

type OrderDto struct {
	OrderId    int    `json:"order_id"`
	Date       string `json:"date"`
	TotalPrice int    `json:"total_price"`
	UserId     int    `json:"user_id"`
}

type OrdersDto []OrderDto

type OrderWithDetailsDto struct {
	OrderId    int        `json:"order_id"`
	Date       string     `json:"date"`
	TotalPrice int        `json:"total_price"`
	UserId     int        `json:"user_id"`
	Details    DetailsDto `json:"details"`
}
