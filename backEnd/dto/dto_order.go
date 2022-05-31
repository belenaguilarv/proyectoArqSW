package dto

type OrderDto struct {
	OrderId int `json:"order_id"`
	//	OrderDate  date `json:"order_date"` // date no lo reconoce
	TotalPrice int `json:"total_price"`
	Quantity   int `json:"quantity"`

	// 	ProductId        int     `json:"product_id"`
}

type OrdersDto []OrderDto
