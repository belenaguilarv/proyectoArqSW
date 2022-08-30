package controllers

import (
	"net/http"
	"strconv"

	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"

	service "github.com/belenaguilarv/proyectoArqSW/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var orderDto dto.OrderDto

	orderDto, err := service.OrderService.GetOrderById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func GetOrders(c *gin.Context) {
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func GetOrdersWithDetails(c *gin.Context) {
	var ordersWithDetailsDTO dto.OrdersWithDetailsDto
	ordersWithDetailsDTO, err := service.OrderService.GetOrdersWithDetails()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersWithDetailsDTO)
}

func GetOrderWithDetailsById(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var orderWithDetailsdto dto.OrderWithDetailsDto
	orderWithDetailsdto, err := service.OrderService.GetOrderWithDetailsById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, orderWithDetailsdto)

}

func InsertOrder(c *gin.Context) {
	var orderDto dto.OrderWithDetailsDto
	err := c.BindJSON(&orderDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, orderDto)
}
