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
	router.GET("/user/:id", userController.GetUserById) // TODO OK
	router.GET("/user", userController.GetUsers)        // TODO OK

	// Login Mapping
	router.POST("/login", userController.LoginUser) //TODO OK

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById) // TODO OK
	router.GET("/order", orderController.GetOrders)        // TODO OK

	router.GET("/ordersWithDetails", orderController.GetOrdersWithDetails)       // TODO OK
	router.GET("/orderWithDetails/:id", orderController.GetOrderWithDetailsById) // TODO OK
	router.POST("/neworder", orderController.InsertOrder)                        // TODO OK
	//router.DELETE("/DeleteCarrito/:id", orderController.DeleteOrder)             // FALTA IMPLEMENTAR

	// NOTA: puedo agregar ESTADO en orden para avisar si esta comprado o en estado carrito

	// Detail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById) // TODO OK
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)        // TODO OK
	router.POST("/neworderDetail", orderDetailController.InsertOrderDetail)  // TODO OK

	log.Info("Finishing mappings configurations")
}
