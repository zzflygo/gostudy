package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"log"
)

func GetConfByConsul(addr, prefix string) (config.Config, error) {
	// 新建配置对象
	consulSource := consul.NewSource(
		// 设置consul地址
		consul.WithAddress(addr),
		// 设置前缀
		consul.WithPrefix(prefix),
		// 是否移除前缀
		consul.StripPrefix(true),
	)
	// 配置初始化
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("consul new config failed err", err)
		return nil, err
	}
	// 加载配置
	err = conf.Load(consulSource)
	if err != nil {
		log.Fatal("load consul conf failed err", err)
		return nil, err
	}

	return conf, nil
}
