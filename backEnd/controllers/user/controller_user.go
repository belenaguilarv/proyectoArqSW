package controllers

import (
	"net/http"
	"strconv"

	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"

	service "github.com/belenaguilarv/proyectoArqSW/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

/*
func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto) //agarra el body de request y lo trata de meter dentro de la estructura dto

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		panic("")
	}

	c.JSON(http.StatusAccepted, tokenDto)

}
*/
