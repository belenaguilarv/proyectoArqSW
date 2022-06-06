package services

import (
	OrderCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients" // conecta a la base de datos
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors" // que hace esto? (utils)
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrdersByUserId(userId int) (dto.OrdersDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	GetOrderWithDetailsbyId(id int) (dto.OrderDto, dto.DetailsDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}
func (s *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {
	var order model.Order = OrderCliente.GetOrderById(id)
	var orderDto dto.OrderDto
	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	orderDto.OrderId = order.Id
	orderDto.Date = order.Date
	orderDto.TotalPrice = order.TotalPrice
	orderDto.UserId = order.UserId
	return orderDto, nil
}
func (s *orderService) GetOrdersByUserId(userId int) (dto.OrdersDto, e.ApiError) {
	var orders model.Orders = OrderCliente.GetOrdersByUserId(userId)
	var ordersDto dto.OrdersDto
	for _, order := range orders {
		var orderDto dto.OrderDto
		orderDto.OrderId = order.Id
		orderDto.Date = order.Date
		orderDto.TotalPrice = order.TotalPrice
		orderDto.UserId = order.UserId
		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil
}
