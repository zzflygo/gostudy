package model

type OrderInfo struct {
	ID           int64          `gorm:"primary_key;auto_increment" json:"id"`
	PayStatus    int32          `json:"pay_status"`
	ShipStatus   int32          `json:"ship_status"`
	OrderPrice   float64        `json:"order_price"`
	OrderDetails []*OrderDetail `gorm:"foreign_key:OrderId"`
}

type OrderDetail struct {
	ID            int64   `gorm:"primary_key;auto_increment" json:"id"`
	OrderId       int64   `gorm:"nut_nil" json:"order_id"`
	ProductId     int64   `json:"product_id"`
	ProductSizeId int64   `json:"product_size_id"`
	ProductNum    int64   `json:"product_num"`
	ProductPrice  float64 `json:"product_price"`
}
