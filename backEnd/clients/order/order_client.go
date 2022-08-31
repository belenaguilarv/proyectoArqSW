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

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)

	log.Debug("Orders: ", orders)

	return orders
}
func UpdateMontoFinal(monto float32, id_Order int) float32 {
	result := Db.Model(&model.Order{}).Where("id = ?", id_Order).Update("monto_final", monto)

	if result.Error != nil {

		log.Error("Order no encontrada")
	}
	return monto
}

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("ERROR al crear la orden")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}

func InsertOrderDetail(order_detail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&order_detail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Detail Created: ", order_detail.Id)
	return order_detail
}
func DeleteOrderById(id int) model.Order {
	var order model.Order
	result := Db.Where("id = ?", id).Delete(&order)

	if result.Error != nil {
		log.Error("Orden no encontrada")
	}

	log.Debug("DELETED order:", order)
	return order

}
func DeleteDetailsByOrderId(orderId int) model.OrderDetails {
	var details model.OrderDetails

	result := Db.Where("order_id", orderId).Delete(&details)

	if result.Error != nil {
		log.Error("error al eliminar los detalles")
	}

	log.Debug("DELETED details:", details)
	return details

}
