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
	_result, _err = fc_open20210406.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	log.Println("args is", args)
	accessKeyId := aliyun.AccessKeyId
	accessKeySecret := aliyun.AccessKeySecret
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	stopStatefulAsyncInvocationHeaders := &fc_open20210406.StopStatefulAsyncInvocationHeaders{}
	stopStatefulAsyncInvocationRequest := &fc_open20210406.StopStatefulAsyncInvocationRequest{}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.StopStatefulAsyncInvocationWithOptions(
		tea.String("xd"),
		tea.String("helloworld"),
		tea.String("taskId"), // 异步任务 id
		stopStatefulAsyncInvocationRequest,
		stopStatefulAsyncInvocationHeaders,
		runtime,
	)
	if _err != nil {
		return _err
	}

	console.Log(util.ToJSONString(tea.ToMap(resp)))

	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
