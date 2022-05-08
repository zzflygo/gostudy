package repositry

import (
	"gorm.io/gorm"
	"order/domain/model"
)

type IOrderRepositry interface {
	CreateTable() error
	GetOrderByID(int64) (*model.OrderInfo, error)
	GetAllOrder() ([]*model.OrderInfo, error)
	CreateOrder(*model.OrderInfo) (int64, error)
	DeleteOrderByID(int64) error
	UpdateShipStatus(*model.OrderInfo) error
	UpdatePayStatus(*model.OrderInfo) error
	UpdateOrderStatus(*model.OrderInfo) error
}

type OrderRepositry struct {
	mysqlDb *gorm.DB
}

func NewOrderRepositry(db *gorm.DB) IOrderRepositry {
	return &OrderRepositry{db}
}

func (o *OrderRepositry) CreateTable() error {
	return o.mysqlDb.AutoMigrate(&model.OrderInfo{}, &model.OrderDetail{})
}

func (o *OrderRepositry) GetOrderByID(id int64) (*model.OrderInfo, error) {
	data := new(model.OrderInfo)
	return data, o.mysqlDb.Model(&model.OrderInfo{}).Preload("order_detail").Where("id=?", id).First(data).Error
}
func (o *OrderRepositry) GetAllOrder() ([]*model.OrderInfo, error) {
	datas := make([]*model.OrderInfo, 100)
	return datas, o.mysqlDb.Model(&model.OrderInfo{}).Find(&datas).Error
}
func (o *OrderRepositry) CreateOrder(orderinfo *model.OrderInfo) (int64, error) {
	return orderinfo.ID, o.mysqlDb.Model(&model.OrderInfo{}).Create(orderinfo).Error
}

func (o *OrderRepositry) DeleteOrderByID(id int64) error {
	tx := o.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Unscoped().Model(&model.OrderInfo{}).Where("id=?", id).Delete(&model.OrderInfo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Model(&model.OrderDetail{}).Where("order_id=?", id).Delete(&model.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
func (o *OrderRepositry) UpdateShipStatus(data *model.OrderInfo) error {
	return o.mysqlDb.Model(&model.OrderInfo{}).Where("id=?", data.ID).UpdateColumn("ship_status", data.ShipStatus).Error
}
func (o *OrderRepositry) UpdatePayStatus(data *model.OrderInfo) error {
	return o.mysqlDb.Model(&model.OrderInfo{}).Where("id=?", data.ID).UpdateColumn("pay_status", data.PayStatus).Error
}
func (o *OrderRepositry) UpdateOrderStatus(data *model.OrderInfo) error {
	return o.mysqlDb.Model(&model.OrderInfo{}).Where("id=?", data.ID).Updates(data).Error
}
