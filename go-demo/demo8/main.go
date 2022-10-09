package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second, // 连接超时时间
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//用于读写 etcd 的键值对
	kv = clientv3.NewKV(client)

	//	create operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "123")

	// exec operation
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("revision:", opResp.Put().Header.Revision)

	getOp = clientv3.OpGet("/cron/jobs/job8")

	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("revision", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("value", string(opResp.Get().Kvs[0].Value))

}
