package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/zzflygo/product/common"
	"github.com/zzflygo/product/proto/product"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
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
	productService := product.NewProductService("go.micro.service.product", src.Client())
	//造数据
	data := NewProcutInfo()

	//run
	res, err := productService.AddProduct(context.TODO(), data)
	if err != nil {
		fmt.Println("ADDproduct 失败了....err:", err)
		fmt.Println(res.Message)
		return
	}
	fmt.Println(res.Id, "+++++", res.Message)
}

func NewProcutInfo() *product.ProductInfo {
	datas := &product.ProductInfo{
		ProductName:        "imooc",
		ProductSku:         "cap",
		ProductPrice:       1.1,
		ProductDescription: "test_message",
		ProductCategoryId:  1,
		ProductImage: []*product.ProductImage{
			{ImageName: "image1",
				ImageUrl:  "url_1",
				ImageCode: "xxx",
			},
			{ImageName: "image2",
				ImageUrl:  "url_2",
				ImageCode: "yyy",
			},
			{ImageName: "image3",
				ImageUrl:  "url_3",
				ImageCode: "zzz",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "体积",
				SizeCode: "平米",
			},
			{SizeName: "重量",
				SizeCode: "kg",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "手机",
			SeoKeywords:    `{"品牌":"小米","内存":"32g"}`,
			SeoDescription: "一个手机",
			SeoCode:        "小米10s",
		},
	}
	return datas
}
