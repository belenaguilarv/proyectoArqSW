package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func GetProductByEan(c *gin.Context) {
	c.JSON(http.StatusOK, "Buscando: "+c.Param("product_id"))
}
*/
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
