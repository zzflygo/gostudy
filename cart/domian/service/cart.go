package service

import (
	"github.com/zzflygo/gostudy/cart/domian/model"
	"github.com/zzflygo/gostudy/cart/domian/repository"
)

type ICartService interface {
	AddCart(info *model.CartInfo) (int64, error)
	DeleteById(int64) error
	ClearCart(int64) error
	Incr(int64, int64) error
	Decr(int64, int64) error
	GetAll(int64) ([]*model.CartInfo, error)
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func NewCartService(cartRepository *repository.CartRepository) ICartService {
	return &CartService{cartRepository}
}

func (c *CartService) AddCart(info *model.CartInfo) (int64, error) {
	return c.CartRepository.AddCart(info)
}
func (c *CartService) DeleteById(UserID int64) error {
	return c.CartRepository.DeleteById(UserID)
}
func (c *CartService) ClearCart(UserID int64) error {
	return c.CartRepository.ClearCart(UserID)
}
func (c *CartService) Incr(CartID int64, num int64) error {
	return c.CartRepository.Incr(CartID, num)
}
func (c *CartService) Decr(CartID int64, num int64) error {
	if num == 1 {
		return c.DeleteById(CartID)
	}
	return c.Decr(CartID, num)
}
func (c *CartService) GetAll(UserID int64) ([]*model.CartInfo, error) {
	return c.CartRepository.GetAll(UserID)
}
