package main

import (
	"context"
	"flag"
	"fmt"
	"gin-reflect-handler/internal/config"
	"gin-reflect-handler/internal/handler"
	"gin-reflect-handler/internal/svc"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `./target - api
	Usage: ./target [-h help] [-c etc/dev.yaml]
	Options:
	`)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&filePath, "c", "etc/dev.yaml", "配置文件所在")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// 配置文件读取
	var c config.Config
	conf.MustLoad(filePath, &c)
	//log.Printf("config is %+v\n", c)

	// 设置logx
	logx.MustSetup(c.Log)

	ctx := svc.NewServiceContext(c)

	address := fmt.Sprintf("%s:%d", ctx.Config.Host, ctx.Config.Port)
	if address == "" {
		logx.Error("[main] can not find any server host config")
		return
	}

	router, err := handler.CreateEngine(ctx)
	if err != nil {
		logx.Error("[main] create engine err: %v", err)
		return
	}

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		if err = server.ListenAndServe(); err != nil {
			if strings.Contains(err.Error(), "bind: address already in use") {
				logx.Error("端口被占用, %s", err.Error())
				panic(err)
				return
			}
			logx.Errorf("[main] start echo server err: %v", err) // 这里不要用 Fatal 不然优雅关停会直接退出
		}
		logx.Infof("[main] echo server is start at: %v", address)
	}()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// DONE: 优雅关停
	for {
		s := <-osSignal
		logx.Infof("[main] 捕获信号 %s", s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			if err = server.Shutdown(ctx); err != nil {
				logx.Errorf("[main] 程序退出异常 err:%s", err.Error())
			} else {
				logx.Info("[main] 程序正常退出")
			}
			logx.Close()
			cancel()
			return
		case syscall.SIGHUP:
			logx.Info("[main] 终端断开信号，忽略")
		default:
			logx.Info("[main] other signal")
		}
	}
}
