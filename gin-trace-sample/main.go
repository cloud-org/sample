package main

import (
	"fmt"
	"gin-trace-sample/middleware"
	"gin-trace-sample/slstrace"
	"gin-trace-sample/trace"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// hello world
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "hello world"})
		return
	})

	// client
	engine.GET("/client", func(ctx *gin.Context) {
		_, span := startSpan(ctx.Request.Context(), "invoke sth")
		defer func() {
			endSpan(span, nil)
		}()
		ctx.JSON(http.StatusOK, gin.H{"client": "ok"})
		return
	})

	// client invoke err
	engine.GET("/clienterr", func(ctx *gin.Context) {
		_, span := startSpan(ctx.Request.Context(), "invoke sth")
		defer func() {
			endSpan(span, fmt.Errorf("invoke error"))
		}()
		ctx.JSON(http.StatusOK, gin.H{"client": "err"})
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
