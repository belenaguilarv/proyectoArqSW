package dto

type OrderDto struct {
	Id         int     `json:"order_id"`
	Date       string  `json:"date"`
	TotalPrice float32 `json:"total_price"`
	UserId     int     `json:"user_id"`
}

type OrdersDto []OrderDto

type OrderWithDetailsDto struct {
<<<<<<< HEAD
	Id         int        `json:"order_id"`
	Date       string     `json:"date"`
	TotalPrice int        `json:"total_price"`
	UserId     int        `json:"user_id"`
	Details    DetailsDto `json:"details"`
=======
	Id         int             `json:"order_id"`
	Date       string          `json:"date"`
	TotalPrice int             `json:"total_price"`
	UserId     int             `json:"user_id"`
	Details    OrderDetailsDto `json:"details"`
>>>>>>> efe4d3285ab3d39a47d34110dcd7bc9c3dc68b13
}
