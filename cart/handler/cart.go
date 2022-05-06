package handler

import (
	"cart/common"
	"cart/domian/model"
	"cart/domian/service"
	"cart/proto/cart"
	"context"
	"log"
)

type Cart struct {
	CartService service.ICartService
}

func (c *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.Response) error {
	cartinfo := new(model.CartInfo)
	err := common.SwapTo(request, cartinfo)
	if err != nil {
		response.Message = "添加到购物车失败"
		return err
	}
	id, err := c.CartService.AddCart(cartinfo)
	if err != nil {
		response.Message = "添加到购物车失败"
		return err
	}
	response.Id = id
	response.Message = "添加成功"
	return nil

}

func (c *Cart) ClearCart(ctx context.Context, request *cart.UserRequest, response *cart.Response) error {

	err := c.CartService.ClearCart(request.UserId)
	if err != nil {
		response.Message = "清空购物车失败"
		return err
	}
	response.Message = "清空购物车成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, request *cart.ChangeNum, response *cart.Response) error {
	err := c.CartService.Incr(request.CartId, request.Num)
	if err != nil {
		response.Message = "添加失败"
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, request *cart.ChangeNum, response *cart.Response) error {
	err := c.CartService.Decr(request.CartId, request.Num)
	if err != nil {
		response.Message = "减少失败"
		return err
	}
	response.Message = "减少成功"
	return nil
}
func (c *Cart) GetAll(ctx context.Context, request *cart.UserRequest, response *cart.CartAllResponse) error {
	infos, err := c.CartService.GetAll(request.UserId)
	if err != nil {
		return err
	}
	for _, v := range infos {
		cartinfo := new(cart.CartInfo)
		if err := common.SwapTo(v, cartinfo); err != nil {
			log.Fatal("get all carts failed err:", err)
			return err
		}
		response.Carts = append(response.Carts, cartinfo)
	}
	return nil
}

func (c *Cart) DeleteByID(ctx context.Context, request *cart.DeleteByIdRequest, response *cart.Response) error {
	err := c.CartService.DeleteById(request.CartId)
	if err != nil {
		response.Message = "删除失败"
		return err
	}
	response.Message = "删除成功"
	return nil
}
