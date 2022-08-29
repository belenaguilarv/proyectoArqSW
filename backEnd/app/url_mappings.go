package app

import (
	orderDetailController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/orderDetail"
	productController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/product"

	userController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/user"

	orderController "github.com/belenaguilarv/proyectoArqSW/backEnd/controllers/order"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById) // TODO OK
	router.GET("/product", productController.GetProducts)        // TODO OK

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById) // No trae username  - pedir ayuda al profe
	router.GET("/user", userController.GetUsers)        // No trae username  - pedir ayuda al profe

	// Login Mapping
	router.POST("/login", userController.LoginUser) //TODO OK

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById) // TODO OK
	router.GET("/order", orderController.GetOrders)        // TODO OK
	//router.GET("/orderWithDetails", orderController.GetOrderWithDetails)         // FALTA IMPLEMENTAR
	//router.GET("/orderWithDetails/:id", orderController.GetOrderWithDetailsById) // FALTA IMPLEMENTAR
	//router.POST("/neworder", orderController.InsertOrderWithDetails)             // anda a medias, revisar
	//router.DELETE("/DeleteCarrito/:id", orderController.DeleteOrder)             // FALTA IMPLEMENTAR

	// Detail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById) //np busca bien por id
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)        // no busca bien el order id
	router.POST("/neworderDetail", orderDetailController.InsertOrderDetail)  // no devuelve el DetailId bien(si lo crea en la BD, pero no lo muestra bien)

	log.Info("Finishing mappings configurations")
}
