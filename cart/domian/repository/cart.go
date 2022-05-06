package repository

import (
	"github.com/zzflygo/gostudy/cart/domian/model"
	"gorm.io/gorm"
)

type ICartRepository interface {
	InitCartTable() error
	AddCart(info *model.CartInfo) (int64, error)
	DeleteById(int64) error
	ClearCart(int64) error
	Incr(int64, int64) error
	Decr(int64, int64) error
	GetAll(int64) ([]*model.CartInfo, error)
}

type CartRepository struct {
	MysqlDB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}

func (c *CartRepository) InitCartTable() error {
	return c.MysqlDB.AutoMigrate(&model.CartInfo{})
}

func (c *CartRepository) AddCart(info *model.CartInfo) (int64, error) {
	n1 := &model.CartInfo{}
	return info.ID, c.MysqlDB.Model(&model.CartInfo{}).FirstOrCreate(n1, info).Error
}

func (c *CartRepository) DeleteById(CartID int64) error {
	return c.MysqlDB.Model(&model.CartInfo{}).Where("id=?", CartID).Delete(&model.CartInfo{}).Error
}
func (c *CartRepository) ClearCart(UserID int64) error {
	return c.MysqlDB.Model(&model.CartInfo{}).Where("user_id=?", UserID).Delete(&model.CartInfo{}).Error
}
func (c *CartRepository) Incr(CartID int64, num int64) error {
	return c.MysqlDB.Model(&model.CartInfo{}).Where("id=?", CartID).UpdateColumn("num", gorm.Expr("num+?", num)).Error
}
func (c *CartRepository) Decr(CartID int64, num int64) error {

	return c.MysqlDB.Model(&model.CartInfo{}).Where("id=?", CartID).UpdateColumn("num", gorm.Expr("num-?", num)).Error
}
func (c *CartRepository) GetAll(UserID int64) ([]*model.CartInfo, error) {
	Cartinfos := make([]*model.CartInfo, 20)
	return Cartinfos, c.MysqlDB.Model(&model.CartInfo{}).Where("user_id=?", UserID).Find(Cartinfos).Error
}
