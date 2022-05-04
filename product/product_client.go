package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"product/common"
)

func main() {

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	//链路追踪
	tc, io, err := common.TraceInit("go.micro.product.client", "127.0.0.1:6831")
	if err != nil {
		fmt.Printf("TraceInit failed err:", err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tc)
	//新建micro服务
	src := micro.NewService(
		//名字
		micro.Name("go.micro.product.client"),
		//版本
		micro.Version("latest"),
		//链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//暴露端口
		micro.Address("127.0.0.1:9192"),
		//添加注册中心
		micro.Registry(consulRegistry),
	)
	//访问服务端..服务
	go_micro_server_product
	//run
}

func SendMessage() {

}
