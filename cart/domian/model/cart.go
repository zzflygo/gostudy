package model

type CartInfo struct {
	ID        int64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID    int64 `gorm:"not_nil" json:"user_id"`
	ProductID int64 `gorm:"not_nil" json:"product_id"`
	SizeID    int64 `gorm:"not_nil" json:"size_id"`
	Num       int64 `gorm:"not_nil" json:"num"`
}
