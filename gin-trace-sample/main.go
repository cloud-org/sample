package main

import (
	"gin-trace-sample/middleware"
	"gin-trace-sample/trace"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	trace.StartAgent(trace.Config{
		Name:     "main",
		Endpoint: "http://localhost:14268/api/traces",
		Sampler:  1.0,
		Batcher:  "jaeger",
	})
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(middleware.TracingHandler("main"))
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "hello world"})
		return
	})
	engine.Run("0.0.0.0:9091")
}
