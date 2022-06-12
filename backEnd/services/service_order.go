package services

import (
	orderCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/order" // conecta a la base de datos
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors" // que hace esto? (utils)
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrdersByUserId(userId int) (dto.OrdersDto, e.ApiError)
	GetOrderDetailsByOrderId(orderId int) (dto.OrderDetailsDto, e.ApiError)
	GetOrdersWithDetailsByUserId(userId int) (dto.OrdersWithDetailsDto, e.ApiError)
	GetOrderWithDetailsbyOrderId(orderId int) (dto.OrderWithDetailsDto, e.ApiError)
	PostOrder()
	DeleteOrder(id int)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {
	var order model.Order = orderCliente.GetOrderById(id)
	var orderDto dto.OrderDto
	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	orderDto.Id = order.Id
	orderDto.Date = order.Date
	orderDto.TotalPrice = order.TotalPrice
	orderDto.Id = order.Id
	return orderDto, nil
}
func (s *orderService) GetOrdersByUserId(userId int) (dto.OrdersDto, e.ApiError) {
	var orders model.Orders = orderCliente.GetOrdersByUserId(userId)
	var ordersDto dto.OrdersDto
	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.Id = order.Id
		orderDto.Date = order.Date
		orderDto.TotalPrice = order.TotalPrice
		orderDto.UserId = order.UserId
		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil
}

func (s *orderService) GetOrdersWithDetailsByUserId(userId int) (dto.OrdersWithDetailsDto, e.ApiError) {
	var ordersWithDetails model.OrdersWithDetails = orderCliente.GetOrdersWithDetailsByUserId(userId)
	var ordersWithDetailsDto dto.OrdersWithDetailsDto

	for _, orderWithDetails := range ordersWithDetails {
		for _, order := range orderWithDetails.Order {
			var orderDto dto.OrderDto
			orderDto.Id = ordersWithDetails.Order.Id
			orderDto.Date = ordersWithDetails.Order.Date
			orderDto.TotalPrice = ordersWithDetails.Order.TotalPrice
			orderDto.UserId = ordersWithDetails.Order.UserId
			var i int = 0
			for _, detail := range orderWithDetails.Details {
				var detailDto dto.DetailDto
				detailDto.Id = ordersWithDetails.Details[i].Id
				detailDto.Price = ordersWithDetails.Details[i].Price
				detailDto.Quantity = ordersWithDetails.Details[i].Quantity
				detailDto.OrderId = ordersWithDetails.Details[i].OrderId
				detailDto.ProductId = ordersWithDetails.Details[i].ProductId
				orderDto.Details = append(orderDto.Details, detailDto)
				i++
			}
			ordersWithDetailsDto = append(ordersWithDetailsDto, orderDto)
		}
		ordersWithDetailsDto = append(ordersWithDetailsDto, orderWithDetails)
	}
	return ordersWithDetailsDto, nil
}

func GetOrderDetailsByOrderId(orderId int) (dto.OrderDetailsDto, e.ApiError) {
	var orderDetails model.OrderDetails = orderCliente.GetOrderDetailsByOrderId(orderId)
	var orderDetailsDto dto.OrderDetailsDto
	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.ProductId = orderDetail.ProductId
		orderDetailDto.OrderId = orderDetail.OrderId
		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}
	return orderDetailsDto, nil
}

func (s *orderService) GetOrderWithDetailsbyOrderId(id int) (dto.OrderWithDetailsDto, e.ApiError) {
	var orderWithDetails model.OrderWithDetails = orderCliente.GetOrderWithDetailsbyOrderId(id)
	var orderWithDetailsDto dto.OrderWithDetailsDto

	for _, detail := range orderWithDetails.Details {
		var detailDto dto.OrderDetailDto
		detailDto.Id = detail.Id
		detailDto.Quantity = detail.Quantity
		detailDto.Price = detail.Price
		detailDto.ProductId = detail.ProductId
		detailDto.OrderId = detail.OrderId
		orderWithDetailsDto.Details = append(orderWithDetailsDto.Details, detailDto)
	}
	orderWithDetailsDto.order = orderWithDetails.order
	return orderWithDetailsDto, nil
}

func (s *orderService) PostOrder() {
	orderCliente.PostOrder() //delega el trabajo al cliente

}
func (s *orderService) DeleteOrder(id int) {
	orderCliente.DeleteOrder(id) //delega el trabajo al cliente
}
