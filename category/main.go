package main

import (
	"category/common"
	"category/domian/repository"
	"category/domian/service"
	"category/handler"
	"category/proto/category"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//consul配置中心
	conf, err := common.GetConfByConsul("127.0.0.1:8500", "micro/config")
	if err != nil {
		log.Fatal("get config form consul failed", err)
		return
	}
	fmt.Println("conf 配置")
	fmt.Println(conf)
	//consul注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{"127.0.0.1:8500"}
	})

	//新建服务对象
	src := micro.NewService(
		micro.Name("category"),
		micro.Version("latest"),
		//暴露的端口
		micro.Address("127.0.0.1:9191"),
		//注册中心
		micro.Registry(consulRegistry),
	)
	//初始化微服务
	src.Init()
	//从配置中心拿到mysql配置
	mysqlstr, err := common.EnCodeMysqlConf(conf, "mysql")
	fmt.Println(mysqlstr)
	//链接gorm
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: mysqlstr}))
	cr := repository.NewCategoryRepository(db)
	//...初始化表格
	err = cr.InitCategory()
	if err != nil {
		return
	}
	err = category.RegisterCategoryHandler(src.Server(), &handler.Category{service.NewCategoryService(cr)})
	if err != nil {
		return
	}
	err = src.Run()
	if err != nil {
		return
	}
}
