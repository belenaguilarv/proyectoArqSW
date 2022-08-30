package services

import (
	orderCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/order"
	orderDetailCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/orderDetail"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderwithdetailsDto dto.OrderWithDetailsDto) (dto.OrderWithDetailsDto, e.ApiError)
	GetOrdersWithDetails() (dto.OrdersWithDetailsDto, e.ApiError)
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
	orderDto.UserId = order.UserId

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrders()
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

func (s *orderService) InsertOrder(orderwithdetailsDto dto.OrderWithDetailsDto) (dto.OrderWithDetailsDto, e.ApiError) {
	var order model.Order
	var details model.OrderDetails

	order.Date = orderwithdetailsDto.Date
	order.TotalPrice = orderwithdetailsDto.TotalPrice
	order.UserId = orderwithdetailsDto.UserId
	order.Id = orderwithdetailsDto.Id

	for _, OrderDetailDto := range orderwithdetailsDto.Details {
		var detail model.OrderDetail
		detail.Id = OrderDetailDto.Id
		detail.Quantity = OrderDetailDto.Quantity
		detail.Price = OrderDetailDto.Price
		detail.TotalPrice = OrderDetailDto.TotalPrice
		detail.ProductId = OrderDetailDto.ProductId
		detail.OrderId = OrderDetailDto.OrderId

		details = append(details, detail)
	}
	orderCliente.InsertOrder(order) // arreglar para el final
	orderCliente.InsertOrderDetails(details)

	return orderwithdetailsDto, nil
}

func (s *orderService) GetOrdersWithDetails() (dto.OrdersWithDetailsDto, e.ApiError) {
	var orders model.Orders = orderCliente.GetOrders()
	var details model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var ordersWithDetailsDto dto.OrdersWithDetailsDto

	for _, order := range orders {
		var orderWithDetailsDto dto.OrderWithDetailsDto

		orderWithDetailsDto.Id = order.Id
		orderWithDetailsDto.Date = order.Date
		orderWithDetailsDto.TotalPrice = order.TotalPrice
		orderWithDetailsDto.UserId = order.UserId

		ordersWithDetailsDto = append(ordersWithDetailsDto, orderWithDetailsDto)
	}

	for j := 0; j < len(ordersWithDetailsDto); j++ {
		var control bool = false
		var detailsssDto dto.OrderDetailsDto

		for i := 0; i < len(details); i++ {
			if details[i].OrderId == ordersWithDetailsDto[j].Id {
				control = true
				var detailDto dto.OrderDetailDto

				detailDto.Id = details[i].Id
				detailDto.Quantity = details[i].Quantity
				detailDto.Price = details[i].Price
				detailDto.TotalPrice = details[i].TotalPrice
				detailDto.ProductId = details[i].ProductId
				detailDto.OrderId = details[i].OrderId

				detailsssDto = append(detailsssDto, detailDto)
			}
		}
		if control {
			ordersWithDetailsDto[j].Details = detailsssDto
		}
	}

	return ordersWithDetailsDto, nil

}
