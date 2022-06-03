package dto

type ProductDto struct {
	ProductId int     `json:"product_id"`
	Name      string  `json:"name"`
	Picture   string  `json:"picture_url"`
	Price     float32 `json:"product_unit_price"`
}

type ProductsDto []ProductDto
