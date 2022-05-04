package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domian/model"
	"user/domian/repository"
)

//control
// 创建service层接口

type IUserDataService interface {
	//增
	AddUser(user *model.User) (int64, error)
	//删
	DeleteUser(int64) error
	//改
	UpdateUser(user *model.User) error
	//查
	FindUserByName(string) (*model.User, error)
	//验证
	CheckPassword(username, password string) (bool, error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

//加密用户密码
func GeneratePassword(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

//对比用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

//确认账号密码是否正确
func (u *UserDataService) CheckPassword(username, password string) (bool, error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return ValidatePassword(password, user.HashPassword)
}

//新建用户
func (u *UserDataService) AddUser(user *model.User) (userid int64, err error) {
	pwd, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.Id, err
	}
	user.HashPassword = string(pwd)
	return u.UserRepository.CreateUser(user)
}

//通过用户id删除用户
func (u *UserDataService) DeleteUser(userid int64) (err error) {
	return u.UserRepository.DeleteUserById(userid)
}

//更新用户信息
func (u *UserDataService) UpdateUser(user *model.User) (err error) {
	return u.UserRepository.UpdateUser(user)
}

//通过username查找用户
func (u *UserDataService) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return u.UserRepository.FindUserByName(name)
}
