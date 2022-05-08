package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil); err != nil {
			log.Fatal("监控启动失败")
			return
		}
		log.Info("监控成功,port:", port)
	}()

}
