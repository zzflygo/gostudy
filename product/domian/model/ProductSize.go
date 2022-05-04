package model

type ProductSize struct {
	ID            int64  `gorm:"primary_key;auto_increment" json:"id"`
	ProductSizeID int64  `gorm:"unique_index" json:"product_size_id"`
	SizeName      string `gorm:"not_nil" json:"size_name"`
	SizeCode      string `gorm:"unique_index" json:"size_code"`
}
