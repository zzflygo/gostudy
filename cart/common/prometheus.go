package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-micro.dev/v4/util/log"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	//prometheus的默认方法
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil); err != nil {
			log.Fatal("监控失败")
			return
		}
		log.Info("成功启动监听,端口:", port)
	}()
}
