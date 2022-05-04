package handler

import (
	"context"
	"user/domian/model"
	"user/domian/service"
	"user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//注册
func (u *User) Register(ctx context.Context, userRes *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRes.UserName,
		FirstName:    userRes.FirstName,
		HashPassword: userRes.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		response.Message = "注册失败"
		return err
	}
	response.Message = "添加成功"
	return nil
}

// Login 登录 每个对外暴露的接口 都有自己装有传参
func (u *User) Login(ctx context.Context, useRes *user.LoginRequest, response *user.LoginResponse) error {
	//验证密码
	ok, err := u.UserDataService.CheckPassword(useRes.UserName, useRes.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = ok
	return nil
}

//查询用户信息
func (u *User) GetUserInfo(ctx context.Context, useRes *user.UserInfoRequest, response *user.UserInfoResponse) (err error) {

	u1, err := u.UserDataService.FindUserByName(useRes.UserName)
	if err != nil {
		return err
	}
	response.UserId = u1.Id
	response.UserName = u1.UserName
	response.FirstName = u1.FirstName
	return nil
}
