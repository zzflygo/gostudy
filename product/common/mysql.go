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

func GetMysqlFromConsul(conf config.Config, path ...string) (string, error) {
	mconf := new(MysqlConfig)
	err := conf.Get(path...).Scan(mconf)
	if err != nil {
		return "", err
	}
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mconf.User, mconf.Pwd, mconf.Host, mconf.Port, mconf.Database)
	return addr, nil
}
