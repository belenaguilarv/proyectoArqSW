package app

import (
	productController "go/controllers/product"
	userController "go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductByEan)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	log.Info("Finishing mappings configurations")
}
