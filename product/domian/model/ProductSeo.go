package model

type ProductSeo struct {
	ID             int64  `gorm:"primary_key;auto_increment;" json:"id"`
	ProductSeoID   int64  `gorm:"unique_index" json:"product_seo_id"`
	SeoTitle       string `json:"seo_title"`
	SeoKeywords    string `json:"seo_keywords"`
	SeoDescription string `json:"seo_description"`
	SeoCode        string `gorm:"not_nil;unique_index" json:"seo_code"`
}
