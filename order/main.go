package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v4"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/zzflygo/gostudy/cart/common"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	common2 "order/common"
	"order/domain/repositry"
	"order/domain/service"
	"order/handler"
	"order/proto/order"
)

func main() {
	//配置中心
	conf, err := common.NewConsulConfig("127.0.0.1:8500", "micro/config")
	if err != nil {
		log.Fatal("NewConsulConfig failed err:", err)
		return
	}
	addr, err := common.GetMysqlConfigFormConsul(conf, "mysql")
	if err != nil {
		log.Fatal("GetMysqlConfigFormConsul failed err:", err)
		return
	}
	fmt.Println(addr)
	//注册中心
	consul2 := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	//链路追踪
	tc, io, err := common.NewTracing("go.micro.service.order")
	if err != nil {
		log.Fatal("NewTracingg failed err:", err)
		return
	}
	opentracing.SetGlobalTracer(tc)
	defer io.Close()
	//监控
	common2.PrometheusBoot(9092)
	//链接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: addr}))
	if err != nil {
		log.Fatal("gorm.Open failed err:", err)
		return
	}
	//设置连接池
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	orderRepositry := repositry.NewOrderRepositry(db)
	err = orderRepositry.CreateTable()
	if err != nil {
		log.Fatal("CreateTable failed err:", err)
		return
	}
	//新建micro-service
	src := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Address("0.0.0.0:9097"),
		micro.Version("latest"),
		//注册
		micro.Registry(consul2),
		//链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(100)),
		//监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	//init
	src.Init()
	//绑定方法
	ods := service.NewOrderService(orderRepositry)
	err = order.RegisterOrderHandler(src.Server(), &handler.Order{OrderService: ods})
	if err != nil {
		log.Fatal("RegisterOrderHandler failed err:", err)
		return
	}
	//run
	src.Run()
}
