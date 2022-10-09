// +build linux

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func main() {
	fc.Start(HandlerData)
}

type Req struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type Resp struct {
	NameAge string `json:"nameAge"`
}

func HandlerData(ctx context.Context, req *Req) (*Resp, error) {
	// 业务处理逻辑
	nameAge := fmt.Sprintf("%v%v", req.Name, req.Age)
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("获取 ctx 失败")
	}
	fctx.GetLogger().Infof("认证 %+v", fctx.Credentials)
	fctx.GetLogger().Infof("服务配置 %+v", fctx.Service)
	time.Sleep(10 * time.Second)
	fctx.GetLogger().Info("休眠 10s")
	return &Resp{
		NameAge: nameAge,
	}, nil
}
