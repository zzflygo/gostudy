package repository

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/domian/model"
)

// 相当于dao层 直接操作数据库
// UserRepository  封装数据库操作对象
type UserRepository struct {
	mysqlDb *gorm.DB
}

//封装接口
type IUserRepository interface {
	//初始化表

	//根据用户名查找用户信息
	FindUserByName(name string) (*model.User, error)
	//根据用户ID找到用户信息
	FindUserById(id int64) (*model.User, error)
	//创建用户
	CreateUser(user *model.User) (userID int64, err error)
	//根据ID删除用户
	DeleteUserById(id int64) error
	//更新用户信息
	UpdateUser(user *model.User) error
	//查找所有用户
	FindAll() ([]*model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

// 初始化表
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})

}

//根据用户名查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name=?", name).Find(user).Error
}

//根据用户ID找到用户信息
func (u *UserRepository) FindUserById(id int64) (user *model.User, err error) {
	user = &model.User{}

	return user, u.mysqlDb.Where("id=?", id).Find(user).Error
}

//创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.Id, u.mysqlDb.Create(user).Error
}

//根据ID删除用户
func (u *UserRepository) DeleteUserById(id int64) error {
	return u.mysqlDb.Where("id=?", id).Delete(&model.User{}).Error
}

//更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Updates(user).Error
}

//查找所有用户
func (u *UserRepository) FindAll() (userAll []*model.User, err error) {
	userAll = make([]*model.User, 1000)
	return userAll, u.mysqlDb.Find(userAll).Error
}
