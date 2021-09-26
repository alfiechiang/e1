package main

import (
	"e1/src/router"
	"e1/start"
	"flag"
	"fmt"

	"e1/global"

	"github.com/spf13/viper"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

var (
	config string
)

func init() {

	start.SetupDBEngine()
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	viper.UnmarshalKey("Database", &global.DBConfig)

}

func main() {
	//app.CreateToken(123)
	trac()
	router.Route()
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("say-hello")
	helloStr := "say-hello"
	println(helloStr)
	span.Finish()

}

func trac() {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "your_service_name",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
}
