package dto

type ProductDto struct {
	Id          int     `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture_url"`
	Price       float32 `json:"product_unit_price"`
	CategoryIds []int   `json:"category_ids"`
}

type ProductsDto []ProductDto
