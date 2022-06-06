package model

type CatXProd struct {
	CategoryId int `json:"category_id"`
	ProductId  int `json:"product_id"`
}

type CatsXProds []CatXProd
