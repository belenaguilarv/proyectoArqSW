package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProducts(c *gin.Context) {

}
func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))
}
func GetProductByCategory(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
