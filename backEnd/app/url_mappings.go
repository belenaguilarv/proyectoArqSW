package app

import (
	orderDetailController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/orderDetail"
	productController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/product"

	userController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/user"

	orderController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/order"

	detailController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/orderDetail"

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

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/order", orderController.GetOrders)
	//router.POST("/neworder", orderController.InsertOrder)

	// Detail Mapping
	router.GET("/orderDetail/:id", detailController.GetOrderDetailById)
	router.GET("/orderDetail", detailController.GetOrderDetails)
	router.POST("/neworderDetail", orderDetailController.InsertOrderDetail)

	log.Info("Finishing mappings configurations")
}
