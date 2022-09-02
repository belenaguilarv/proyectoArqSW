package controllers

import (
	"net/http"
	"strconv"

	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"

	service "github.com/belenaguilarv/proyectoArqSW/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProductsBYpalabra(c *gin.Context) {
	var palabraClave dto.PalabraClaveDto
	err := c.BindJSON(&palabraClave)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productsDto, er := service.ProductService.GetProductsByPalabrasClaves(palabraClave.Clave)

	if er != nil {
		c.JSON(er.Status(), er)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, productsDto)

}
