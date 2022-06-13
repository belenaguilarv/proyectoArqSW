package app

import (
	productController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/product"

	userController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/user"

	orderController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/order"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)

	// Login Mapping
	router.POST("/login", userController.LoginUser)

	// Order Mapping
	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/order", orderController.GetOrders)

	log.Info("Finishing mappings configurations")
}
