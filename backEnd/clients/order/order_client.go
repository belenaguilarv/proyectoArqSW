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

/*
func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)
	log.Debug("Orders: ", orders)
	return orders
}


func GetOrderWithDetailsbyId(id int) model.OrderWithDetails {
	var order model.OrderWithDetails
	Db.Where("id = ?", id).First(&order)
	log.Debug("Order: ", order)
	return order
}
*/
func PostOrder() {
	// necesito crear los detalles asociados a la orden
	//Db.Create(&model.OrderWithDetails{}.Details)
	Db.Create(&model.Order{})
}
func DeleteOrder(id int) {
	// necesito borrar los detalles asociados a la orden
	Db.Where("order_id = ?", id).Find(&model.OrderDetails{}).Delete(&model.OrderDetails{})
	Db.Where("id = ?", id).Delete(&model.Order{})
}
