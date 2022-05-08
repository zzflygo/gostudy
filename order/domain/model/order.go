package service

type OrderInfo struct {
	ID           int64          `gorm:"primary_key;auto_increment" json:"id"`
	PayStatus    int32          `json:"pay_status"`
	ShipStatus   int32          `json:"ship_status"`
	OrderPrice   float64        `json:"order_price"`
	OrderDetails []*OrderDetail `gorm:"foreign_key:OrderDetail:order_id"`
}

type OrderDetail struct {
	ID            int64
	OrderId       int64
	ProductId     int64
	ProductSizeId int64
	ProductNum    int64
	ProductPrice  float64
}
