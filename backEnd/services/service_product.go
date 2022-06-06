package services

import (
	ProductCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients" // conecta a la base de datos
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors" // que hace esto? (utils)
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProductByCategory(categoryId int) (dto.ProductsDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}
func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {
	var product model.Product = ProductCliente.GetProductById(id)
	var productDto dto.ProductDto
	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.Name = product.Name
	productDto.Id = product.Id
	productDto.Price = product.Price
	productDto.CategoryId = product.CategoryId
	return productDto, nil
}
func (s *productService) GetProductByCategory(categoryId int) (dto.ProductsDto, e.ApiError) {
	var products model.Products = ProductCliente.GetProductByCategory(categoryId)
	var productsDto dto.ProductsDto
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Id = product.Id
		productDto.Price = product.Price
		productDto.CategoryId = product.CategoryId
		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}
func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {
	var products model.Products = ProductCliente.GetProducts()
	var productsDto dto.ProductsDto
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Id = product.Id
		productDto.Price = product.Price
		productDto.CategoryId = product.CategoryId
		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil
}
