package services

import (
	productCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/product"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}
func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
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

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {
	var products model.Products = productCliente.GetProducts()
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
