package order

import (
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order
	Db.Where("id = ?", id).First(&order)
	log.Debug("Order: ", order)
	return order
}

func GetOrdersByUserId(userId int) model.Orders {
	var orders model.Orders
	Db.Where("user_id = ?", userId).Find(&orders)
	log.Debug("Orders: ", orders)
	return orders
}

func GetOrderDetailsByOrderId(orderId int) model.OrderDetails {
	var details model.OrderDetails
	Db.Where("order_id = ?", orderId).Find(&details)
	log.Debug("Order Details: ", details)
	return details
}

func GetOrdersWithDetailsByUserId(userId int) model.OrdersWithDetails {
	var ordersWithDetails model.OrdersWithDetails
	var orders model.Orders
	var details model.OrderDetail

	orders = GetOrdersByUserId(userId)
	for _, order := range orders {
		details = GetOrderDetailsByOrderId(order.Id)
		ordersWithDetails = append(ordersWithDetails, model.OrdersWithDetails{Order: order, OrderDetails: details})
	}
	log.Debug("Orders With Details: ", ordersWithDetails)
	return ordersWithDetails

}

func GetOrderWithDetailsbyOrderId(id int) model.OrderWithDetails {
	var OrderWithDetails model.OrderWithDetails
	var order model.Orders
	var details model.OrderDetail

	order = GetOrderById(id)
	for _, order := range order {
		details = GetOrderDetailsByOrderId(order.Id)
		OrderWithDetails = append(OrderWithDetails, model.OrderWithDetails{Order: order, OrderDetails: details})
	}
	log.Debug("Order With Details: ", OrderWithDetails)
	return OrderWithDetails
}

func PostOrder() {
	// necesito crear los detalles asociados a la orden
	var orderWithDetails model.OrderWithDetails = GetOrderWithDetailsbyOrderId(id)
	var order model.Order = orderWithDetails.Order
	var details model.OrderDetails = orderWithDetails.Details

	for _, detail := range details {
		Db.Create(&detail)
	}
	Db.Create(&order)
}

func DeleteOrder(id int) {
	var OrderWithDetails model.OrderWithDetails = GetOrderWithDetailsbyOrderId(id)
	var order model.Order = OrderWithDetails.Order
	var details model.OrderDetails = OrderWithDetails.details

	for _, detail := range details {
		Db.Delete(&detail)
	}
	Db.Delete(&order)

}
