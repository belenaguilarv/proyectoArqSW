package dto

type ProductDto struct {
	Id          int     `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture_url"`
	Price       float32 `json:"product_unit_price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
}

type ProductsDto []ProductDto
