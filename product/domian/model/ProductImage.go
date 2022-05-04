package model

type ProductImage struct {
	ID             int64  `gorm:"primary_key;auto_increment" json:"id"`
	ProductImageID int64  `json:"product_image_id"`
	ImageName      string `gorm:"nut_nil" json:"image_name"`
	ImageUrl       string `json:"image_url"`
	ImageCode      string `gorm:"not_nil;unique_index" json:"image_code"`
}
