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

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var orderDto dto.OrderDto

	orderDto, err := service.OrderService.GetOrderById(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func GetOrdersByUserId(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var ordersDto dto.OrdersDto

	ordersDto, err := service.OrderService.GetOrdersByUserId(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}

func GetOrderDetailsByOrderId(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var orderDetailsDto dto.OrderDetailsDto

	orderDetailsDto, err := service.OrderService.GetOrderDetailsByOrderId(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailsDto)
}

func GetOrdersWithDetailsByUserId(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var ordersWithDetailsDto dto.OrdersWithDetailsDto

	ordersWithDetailsDto, err := service.OrderService.GetOrdersWithDetailsByUserId(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersWithDetailsDto)
}

func GetOrderWithDetailsbyOrderId(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var orderWithDetailsDto dto.OrderWithDetailsDto

	orderWithDetailsDto, err := service.OrderService.GetOrderWithDetailsbyOrderId(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, orderWithDetailsDto)
}

func Postorder(c *gin.Context) {
	log.Debug("Post order")
	var orderDto dto.OrderDto
	err := c.ShouldBindJSON(&orderDto) //se recibe el json y se pasa a la estructura
	errr := service.OrderService.PostOrder(orderDto)

	if err != nil || errr != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func DeleteOrder(c *gin.Context) {
	log.Debug("Delete order")

	id, _ := strconv.Atoi(c.Param("id"))        //se pasa el id de array a int
	err := service.OrderService.DeleteOrder(id) //delega el trabajo al service

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
