package model

//定义模型

type User struct {
	//主键
	Id int64 `gorm:"primary_key;not_null;increment"`
	//用户名称
	UserName string `gorm:"unique_index;not_null"`
	//添加需要的字段
	FirstName string
	//...
	//密码
	HashPassword string
}
