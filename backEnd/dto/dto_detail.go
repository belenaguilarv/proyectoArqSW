package dto

type DetailDto struct {
	DetailId   int `json:"detail_id"`
	Quantity   int `json:"quantity"`
	Price      int `json:"price"`
	TotalPrice int `json:"total_price"`
	ProductId  int `json:"product_id"`
	OrderId    int `json:"order_id"`
}

type DetailsDto []DetailDto
