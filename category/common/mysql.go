package common

import (
	"fmt"
	"go-micro.dev/v4/config"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"point"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func EnCodeMysqlConf(conf config.Config, path ...string) (str string, err error) {
	dbconf := &MysqlConfig{}
	err = conf.Get(path...).Scan(dbconf)
	fmt.Println("mysql的conf", dbconf)
	str = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconf.User, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.Database)
	return
}
