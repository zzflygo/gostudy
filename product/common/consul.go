package common

import (
	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
)

func GetConsulConfig(addr, prefix string) (config.Config, error) {

	consulSource := consul.NewSource(
		//配置consul前缀
		consul.WithPrefix(prefix),
		//配置地址...
		consul.WithAddress(addr),
		consul.StripPrefix(true),
	)
	//配置初始化
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	//加载配置
	err = conf.Load(consulSource)
	if err != nil {
		return nil, err
	}
	return conf, err
}
