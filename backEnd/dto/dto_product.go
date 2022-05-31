package dto

type ProductDto struct {
	ProductId        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
	ProductPicture   string  `json:"product_picture_url"`
	ProductUnitPrice float32 `json:"product_unit_price"`
}

type ProductsDto []ProductDto
