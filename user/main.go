package main

import (
	"fmt"
	"go-micro.dev/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/domian/repository"
	service2 "user/domian/service"
	"user/handler"
	"user/proto/user"
)

func main() {
	// create a new service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	// initialise flags
	srv.Init()
	//初始化服务器db
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:hsp@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local",
	}))
	if err != nil {
		fmt.Println("gorm content mysql failed err:", err)
		return
	}
	fmt.Println("content db success..")
	//建立表格(只用执行一次)
	//err = repository.InitTable(db)
	//if err != nil {
	//fmt.Println("create Table failed err:", err)
	//return
	//}
	// 绑定方法

	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println("RegisterUserHandler failed err:", err)
		return
	}
	// start the service
	if err := srv.Run(); err != nil {
		fmt.Println("start service failed...", err)
		return
	}
}
