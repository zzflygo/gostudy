package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/util/log"
)

func NewConsulConfig(addr, prefix string) (config.Config, error) {
	// 添加配置中心路径
	consulConfig := consul.NewSource(
		consul.WithAddress(addr),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	// 新建配置
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal("consul NewConfig failed err:", err)
		return nil, err
	}
	// 配置加载配置中心文件
	err = conf.Load(consulConfig)
	if err != nil {
		log.Fatal("consul Load NewConfig failed err:", err)
		return nil, err
	}
	return conf, nil
}
