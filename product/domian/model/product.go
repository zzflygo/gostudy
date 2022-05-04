package model

type Product struct {
	ID                 int64           `gorm:"primary_key;auto_increment" json:"id"`
	ProductCategoryID  int64           `json:"product_category_id"`
	ProductName        string          `gorm:"not_nil;unique_index" json:"product_name"`
	ProductSku         string          `gorm:"not_nil;unique_index" json:"product_sku"`
	ProductDescription string          `json:"product_description"`
	ProductPrice       float64         `gorm:"not_nil" json:"product_price"`
	ProductImage       []*ProductImage `gorm:"foreignKey:ProductImageID" json:"product_image"`
	ProductSize        []*ProductSize  `gorm:"foreignKey:ProductSizeID" json:"product_size"`
	ProductSeo         *ProductSeo     `gorm:"foreignKey:ProductSeoID" json:"product_seo"`
}
