// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"fc-in-action/aliyun"
	"log"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	fc_open20210406 "github.com/alibabacloud-go/fc-open-20210406/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *fc_open20210406.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(aliyun.AccountId + aliyun.Endpoint)
	_result = &fc_open20210406.Client{}
	_result, _err = fc_open20210406.NewClient(config)
	return _result, _err
}

type OrderData struct {
	TicketId   string `json:"ticketId"`
	ActivityId string `json:"activityId"`
}

type Req struct {
	Event     string     `json:"event"` // ping or order
	OrderData *OrderData `json:"orderData"`
}

func _main(args []*string) (_err error) {
	log.Println("args is", args)
	accessKeyId := aliyun.AccessKeyId
	accessKeySecret := aliyun.AccessKeySecret
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	invokeFunctionHeaders := &fc_open20210406.InvokeFunctionHeaders{
		XFcInvocationType: tea.String("Async"), // 同步调用
	}
	// 封装请求体 also can json.Marshal struct
	body := []byte(`{"payload": "ping"}`)
	invokeFunctionRequest := &fc_open20210406.InvokeFunctionRequest{
		Body: body,
	}
	runtime := &util.RuntimeOptions{}
	serverName := "xd"         // 服务名
	functionName := "orderrpc" // 函数名
	resp, _err := client.InvokeFunctionWithOptions(tea.String(serverName), tea.String(functionName), invokeFunctionRequest, invokeFunctionHeaders, runtime)
	if _err != nil {
		return _err
	}

	console.Log(tea.String(string(resp.Body)))

	console.Log(util.ToJSONString(tea.ToMap(resp)))
	console.Log(tea.String("异步调用成功"))
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
