package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
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

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("revision", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("Prev is", string(putResp.PrevKv.Value))
		}
	}

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1", clientv3.WithCountOnly()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("res", getResp.Count)
	}

	// 读取前缀为 /cron/jobs/ 的所有 key
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(getResp.Kvs)
	}
}
