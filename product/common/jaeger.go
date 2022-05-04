package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

//创建链路追踪实例
func TraceInit(servicename, addr string) (opentracing.Tracer, io.Closer, error) {
	//创建配置jaeger配置
	cfg := config.Configuration{
		//注册配置名字
		ServiceName: servicename,
		//设置采样器配置
		Sampler: &config.SamplerConfig{
			//
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  addr,
		},
	}
	return cfg.NewTracer()
}
