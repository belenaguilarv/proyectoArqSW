package services

import (
	orderDetailCliente "github.com/belenaguilarv/proyectoArqSW/backEnd/clients/orderDetail"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/dto"
	e "github.com/belenaguilarv/proyectoArqSW/backEnd/errors"
	"github.com/belenaguilarv/proyectoArqSW/backEnd/model"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetails() (dto.OrderDetailsDto, e.ApiError)
	//InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail = orderDetailCliente.GetOrderDetailById(id)
	var orderDetailDto dto.OrderDetailDto

	if orderDetail.DetailId == 0 {
		return orderDetailDto, e.NewBadRequestApiError("orderDetail not found")
	}
	orderDetailDto.DetailId = orderDetail.DetailId
	orderDetailDto.Quantity = orderDetail.Quantity
	orderDetailDto.Price = orderDetail.Price
	orderDetailDto.TotalPrice = int(orderDetail.TotalPrice)
	orderDetailDto.OrderId = orderDetail.OrderId
	orderDetailDto.ProductId = orderDetail.ProductId
	return orderDetailDto, nil
}

//get array
func (s *orderDetailService) GetOrderDetails() (dto.OrderDetailsDto, e.ApiError) {

	var orderDetails model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var orderDetailsDto dto.OrderDetailsDto

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.DetailId = orderDetail.DetailId
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.TotalPrice = int(orderDetail.TotalPrice)
		orderDetailDto.OrderId = orderDetail.OrderId
		orderDetailDto.ProductId = orderDetail.ProductId

		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

/*
func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail
	orderDetail.Cantidad = orderDetailDto.Cantidad
	orderDetail.Precio_Unitario = orderDetailDto.Precio_Unitario
	orderDetail.Total = orderDetailDto.Precio_Unitario * orderDetailDto.Cantidad
	orderDetail.Id_Order = orderDetailDto.Id_Order
	orderDetail.Id_Product = orderDetailDto.Id_Producto

	orderDetail = orderDetailCliente.InsertOrderDetail(orderDetail)

	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Total = orderDetail.Total

	return orderDetailDto, nil
}
*/
