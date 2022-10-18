package slstrace

import (
	"github.com/aliyun-sls/opentelemetry-go-provider-sls/provider"
	"log"
)

type TraceConfig struct {
	ServiceName           string // need unique
	ServiceVersion        string // need unique
	TraceExporterEndpoint string
	Project               string
	InstanceID            string
	AccessKeyID           string
	AccessKeySecret       string
}

//TraceInit sls trace init
func TraceInit(config *TraceConfig) (*provider.Config, error) {
	if config == nil {
		log.Println("nil sls trace config")
		return nil, nil
	}
	slsConfig, err := provider.NewConfig(provider.WithServiceName(config.ServiceName),
		provider.WithServiceVersion(config.ServiceVersion),
		provider.WithTraceExporterEndpoint(config.TraceExporterEndpoint),
		provider.WithMetricExporterEndpoint(""), // not need metrics
		provider.WithSLSConfig(
			config.Project,
			config.InstanceID,
			config.AccessKeyID,
			config.AccessKeySecret,
		))
	// 如果初始化失败则panic，可以替换为其他错误处理方式
	if err != nil {
		log.Printf("provider.NewConfig err: %v", err)
		return nil, err
	}

	if err = provider.Start(slsConfig); err != nil {
		log.Printf("provider.Start err: %v", err)
		return nil, err
	}
	//defer provider.Shutdown(slsConfig)
	return slsConfig, nil
}

func TraceStop(config *provider.Config) {
	if config == nil {
		log.Println("provider sls trace config is nil")
		return
	}
	provider.Shutdown(config)
}
