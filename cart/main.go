package cart

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/zzflygo/gostudy/cart/domian/repository"
	"github.com/zzflygo/gostudy/cart/domian/service"
	"github.com/zzflygo/gostudy/cart/handler"
	"github.com/zzflygo/gostudy/cart/proto/cart"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//连接配置中心
	conf, err := common.NewConsulConfig("127.0.0.1:8500", "micro/config")
	if err != nil {
		fmt.Println("common.NewConsulConfig failed err:", err)
		return
	}
	fmt.Println("连接配置中心成功...")
	//连接注册中心
	consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	fmt.Println("连接注册中心成功...")
	//从配置中心拿到mysql配置
	addr, err := common.GetMysqlConfigFormConsul(conf, "mysql")
	if err != nil {
		fmt.Println("GetMysqlConfigFormConsul failed err:", err)
		return
	}
	fmt.Println("get mysql 配置 成功...")
	//连接mysql
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: addr}))
	if err != nil {
		fmt.Println("gorm.Open failed err:", err)
		return
	}
	fmt.Println("连接mysql成功...")
	Rp := repository.NewCartRepository(db)

	//链路追踪
	tc, io, err := common.NewTracing("go.micro.service.cart")
	if err != nil {
		fmt.Println("common.NewTracing failed err:", err)
		return
	}
	fmt.Println("获取tracing成功...")
	defer io.Close()
	opentracing.InitGlobalTracer(tc)
	//创建micro对象
	src := micro.NewService(
		//服务名字
		micro.Name("go.micro.service.cart"),
		//服务版本
		micro.Version("latest"),
		//服务地址
		micro.Address("127.0.0.1:9200"),
		//链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//qps限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(100)),
	)
	//初始化micro
	src.Init()
	//添加服务
	err = cart.RegisterCartHandler(src.Server(), &handler.Cart{CartService: service.NewCartService(Rp)})
	if err != nil {
		fmt.Println("RegisterCartHandler failed err:", err)
		return
	}
	fmt.Println("注册CartHandler服务成功...")
	//run
	err = src.Run()
	if err != nil {
		fmt.Println("Run failed err:", err)
		return
	}
}
