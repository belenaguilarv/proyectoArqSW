package controllers

import (
	"net/http"
	"strconv"

	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"

	service "github.com/belenaguilarv/proyectoArqSW/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)
jwtKey = []byte("secret_key")
dto.constants.JWT_SECRET_KEY = jwtKey

func GetOrderById(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var orderDto dto.OrderDto

	orderDto, err := service.UserService.GetOrderById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, orderDto)
}


