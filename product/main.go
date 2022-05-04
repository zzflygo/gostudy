package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product/common"
	"product/domian/repository"
	"product/domian/service"
	"product/handler"
	"product/proto/product"
)

func main() {
	// 1.consul配置中心
	config, err := common.GetConsulConfig("127.0.0.1:8500", "micro/config")
	if err != nil {
		log.Fatal("GetConsulConfig failed err:", err)
		return
	}
	fmt.Println("GetConsulConfig success...")
	// 2.consul注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	fmt.Println("consul Registry success...")
	// 3.创建链路追踪jaeger对象
	trace, io, err := common.TraceInit("go.micro.service.product", "127.0.0.1:6831")
	defer io.Close()
	opentracing.SetGlobalTracer(trace)
	// 3.从配置中心拿到mysql配置
	addr, err := common.GetMysqlFromConsul(config, "mysql")
	fmt.Println("get mysql addr success addr:", addr)
	// 4.gorm连接mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: addr,
	}))
	fmt.Println("content mysql success...")
	// 4.1创建Product对象
	prd := repository.NewProductRepository(db)
	// 4.2初始化表
	err = prd.CreateTable()
	// 5.设置micro服务对象
	src := micro.NewService(
		//设置名字
		micro.Name("go.micro.service.product"),
		//设置版本
		micro.Version("latest"),
		//设置服务地址
		micro.Address("127.0.0.1:9100"),
		//设置consul注册中心
		micro.Registry(consulRegistry),
		//设置链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//...
		//设置熔断
		//设置负载均衡
	)
	// 6.init micro对象
	src.Init()
	// 7.绑定方法
	// 7.1
	prdsrc := service.NewProductService(prd)
	err = product.RegisterProductHandler(src.Server(), &handler.Product{ProductData: prdsrc})
	// 8.Run
	src.Run()
}
