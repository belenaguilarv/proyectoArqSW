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



