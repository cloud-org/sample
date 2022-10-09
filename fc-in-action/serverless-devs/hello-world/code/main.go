package main

import (
	"encoding/json"

	gr "github.com/awesome-fc/golang-runtime"
)

func initialize(ctx *gr.FCContext) error {
	ctx.GetLogger().Infoln("init golang!")
	return nil
}

type Req struct {
	Name string `json:"name"`
}

func handler(ctx *gr.FCContext, event []byte) ([]byte, error) {
	fcLogger := gr.GetLogger().WithField("requestId", ctx.RequestID)
	var req Req
	err := json.Unmarshal(event, &req)
	if err != nil {
		fcLogger.Errorf("反序列化失败: %v", err)
		return nil, err
	}
	fcLogger.Infof("hello golang! from hook")
	fcLogger.Infof("name is %v", req.Name)
	return event, nil
}

func main() {
	gr.Start(handler, initialize)
}
