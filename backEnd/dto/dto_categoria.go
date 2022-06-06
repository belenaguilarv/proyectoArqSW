package dto

type categoryDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Products ProductsDto `json:"products"` // podria hacer falta para mejorar la busqueda
}

type categoriesDto []categoryDto
