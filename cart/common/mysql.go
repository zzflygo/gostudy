package common

import (
	"fmt"
	"go-micro.dev/v4/config"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlConfigFormConsul(conf config.Config, path ...string) (addr string, err error) {
	mysqlconf := new(MysqlConfig)
	//此处path为consul配置中心.前缀以后的部分
	err = conf.Get(path...).Scan(mysqlconf)
	if err != nil {
		return "", err
	}
	addr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlconf.User, mysqlconf.Pwd, mysqlconf.Host, mysqlconf.Port, mysqlconf.Database)
	return addr, nil
}
