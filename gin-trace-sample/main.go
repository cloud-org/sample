package main

import (
	"gin-trace-sample/middleware"
	"gin-trace-sample/slstrace"
	"gin-trace-sample/trace"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	trace.StartAgent(trace.Config{
		Name:     "main",
		Endpoint: "http://localhost:14268/api/traces",
		Sampler:  1.0,
		Batcher:  "jaeger",
	})
	//initSlsTrace() // use aliyun sls trace
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(middleware.TracingHandler("main"))
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "hello world"})
		return
	})
	engine.Run("0.0.0.0:9091")
}

func initSlsTrace() {
	_, err := slstrace.TraceInit(&slstrace.TraceConfig{
		ServiceName:           "slstrace-main",
		ServiceVersion:        "v0.0.1-test",
		TraceExporterEndpoint: "",
		Project:               "",
		InstanceID:            "",
		AccessKeyID:           "",
		AccessKeySecret:       "",
	})
	if err != nil {
		log.Printf("sls trace init err: %v", err)
		return
	}
}
