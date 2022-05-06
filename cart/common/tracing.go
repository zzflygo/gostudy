package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracing(sname string) (opentracing.Tracer, io.Closer, error) {
	//设置链路追踪全局
	conf := config.Configuration{
		ServiceName: sname,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort:  "127.0.0.1:6831",
			BufferFlushInterval: time.Second,
			LogSpans:            true,
		},
	}
	return conf.NewTracer()
}
