package model

type Category struct {
	Id          uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Name        string `gorm:"unique_index;nut_nil" json:"name"`
	Level       uint32 `json:"level"`
	Parent      string `json:"parent"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
