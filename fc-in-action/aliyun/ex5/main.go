// +build linux

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"github.com/cloud-org/msgpush"
)

func main() {
	fc.Start(HandlerData)
}

type OrderData struct {
	TicketId   string `json:"ticketId"`
	ActivityId string `json:"activityId"`
}

type Req struct {
	Event     string     `json:"event"` // ping or order
	OrderData *OrderData `json:"orderData"`
}

type Resp struct {
	Event string `json:"event"`
	Rtime string `json:"rtime"`
}

func HandlerData(ctx context.Context, req *Req) (*string, error) {
	// 业务处理逻辑
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("获取 ctx 失败")
	}
	switch req.Event {
	case "ping":
		return tea.String(Ping(fctx)), nil
	case "orderSync":
		OrderLogic(fctx, req)
		return tea.String("order sync invoke"), nil
	case "order":
		fctx.GetLogger().Infof("请求 id %+v", fctx.RequestID)
		go OrderLogic(fctx, req)
		return tea.String("order async invoke"), nil
	default:
		return nil, fmt.Errorf("event 类型错误")
	}
}

func Ping(fctx *fccontext.FcContext) string {
	value := time.Now().Format(time.RFC3339)
	fctx.GetLogger().Info(value)
	return value
}

func OrderLogic(fctx *fccontext.FcContext, req *Req) {
	fctx.GetLogger().Infof("服务配置 %+v", fctx.RequestID)
	time.Sleep(10 * time.Second)
	fctx.GetLogger().Info("休眠 10s")
	// push to pushdeer
	// url := "http://127.0.0.1:9000/?requestId="
	// resp, err := req2.Get(url + fctx.RequestID)
	// if err != nil {
	// 	log.Println("err, ", err)
	// 	return
	// }
	// log.Println("resp is", resp.String())
	mp := msgpush.NewPushDeer("")
	err := mp.Send(fmt.Sprintf("%v-%v", fctx.RequestID, *req.OrderData))
	log.Println("err is", err)
}
