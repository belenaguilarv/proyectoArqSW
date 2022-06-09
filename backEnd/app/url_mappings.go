package app

import (
	productController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/product"

	userController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	//	router.GET("/product/:product_id", productController.GetProductByEan)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	//router.POST("/user", userController.UserInsert)

	router.POST("/login", userController.LoginUser)

	log.Info("Finishing mappings configurations")
}
