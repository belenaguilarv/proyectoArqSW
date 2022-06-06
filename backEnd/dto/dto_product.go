package dto

type ProductDto struct {
	ProductId   int     `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture_url"`
	Price       float32 `json:"product_unit_price"`
	// Categories categoriesDto `json:"categories"` // podria hacer falta para mejorar la busqueda
}

type ProductsDto []ProductDto
