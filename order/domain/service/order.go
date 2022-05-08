package service

import (
	"order/domain/model"
	"order/domain/repositry"
)

type IOrderService interface {
	GetOrderByID(int64) (*model.OrderInfo, error)
	GetAllOrder() ([]*model.OrderInfo, error)
	CreateOrder(*model.OrderInfo) (int64, error)
	DeleteOrderByID(int64) error
	UpdateShipStatus(*model.OrderInfo) error
	UpdatePayStatus(*model.OrderInfo) error
	UpdateOrderStatus(*model.OrderInfo) error
}

type OrderService struct {
	OrderRepositry repositry.IOrderRepositry
}

func NewOrderService(orderRepositry repositry.IOrderRepositry) *OrderService {
	return &OrderService{orderRepositry}
}

func (o *OrderService) GetOrderByID(orderId int64) (*model.OrderInfo, error) {
	return o.OrderRepositry.GetOrderByID(orderId)
}
func (o *OrderService) GetAllOrder() ([]*model.OrderInfo, error) {
	return o.OrderRepositry.GetAllOrder()
}
func (o *OrderService) CreateOrder(data *model.OrderInfo) (int64, error) {
	return o.OrderRepositry.CreateOrder(data)
}
func (o *OrderService) DeleteOrderByID(orderId int64) error {
	return o.OrderRepositry.DeleteOrderByID(orderId)
}
func (o *OrderService) UpdateShipStatus(data *model.OrderInfo) error {
	return o.OrderRepositry.UpdateShipStatus(data)
}
func (o *OrderService) UpdatePayStatus(data *model.OrderInfo) error {
	return o.OrderRepositry.UpdatePayStatus(data)
}
func (o *OrderService) UpdateOrderStatus(data *model.OrderInfo) error {
	return o.OrderRepositry.UpdateOrderStatus(data)
}
