package handler

import (
	"context"
	"fmt"
	"github.com/zzflygo/gostudy/cart/common"
	"order/domain/model"
	"order/domain/service"
	"order/proto/order"
)

type Order struct {
	OrderService *service.OrderService
}

func (o *Order) GetOrderByID(ctx context.Context, req *order.RequestByID, res *order.OrderInfo) error {
	orderinfo, err := o.OrderService.GetOrderByID(req.OrderId)
	if err != nil {
		return err
	}
	err = common.SwapTo(orderinfo, res)
	if err != nil {
		return err
	}
	return nil
}
func (o *Order) GetAllOrder(ctx context.Context, req *order.GetAllRequest, res *order.AllOrderInfo) error {
	orderslice, err := o.OrderService.GetAllOrder()
	if err != nil {
		return err
	}
	for _, info := range orderslice {
		tmp := new(order.OrderInfo)
		err := common.SwapTo(info, tmp)
		if err != nil {
			return err
		}
		res.AllOrderInfo = append(res.AllOrderInfo, tmp)
	}
	return nil
}
func (o *Order) CreateOrder(ctx context.Context, req *order.OrderInfo, res *order.Response) error {
	tmp := new(model.OrderInfo)
	err := common.SwapTo(req, tmp)
	if err != nil {
		res.Msg = "创建失败"
		return err
	}
	id, err := o.OrderService.CreateOrder(tmp)
	if err != nil {
		res.Msg = "创建失败"
		return err
	}

	res.Msg = fmt.Sprintf("创建成功 order_id:%d", id)
	return nil
}
func (o *Order) DeleteOrderByID(ctx context.Context, req *order.RequestByID, res *order.Response) error {
	err := o.OrderService.DeleteOrderByID(req.OrderId)
	if err != nil {
		res.Msg = "删除失败"
		return err
	}
	res.Msg = "删除成功"
	return nil
}
func (o *Order) UpdateShipStatus(ctx context.Context, req *order.UpdateShipRes, res *order.Response) error {
	tmp := new(model.OrderInfo)
	err := common.SwapTo(req, tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	err = o.OrderService.UpdateShipStatus(tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	res.Msg = "更新成功"
	return nil
}
func (o *Order) UpdatePayStatus(ctx context.Context, req *order.UpdatePayRes, res *order.Response) error {
	tmp := new(model.OrderInfo)
	err := common.SwapTo(req, tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	err = o.OrderService.UpdatePayStatus(tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	res.Msg = "更新成功"
	return nil
}
func (o *Order) UpdateOrderStatus(ctx context.Context, req *order.OrderInfo, res *order.Response) error {
	tmp := new(model.OrderInfo)
	err := common.SwapTo(req, tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	err = o.OrderService.UpdateOrderStatus(tmp)
	if err != nil {
		res.Msg = "更新失败"
		return err
	}
	res.Msg = "更新成功"
	return nil
}
