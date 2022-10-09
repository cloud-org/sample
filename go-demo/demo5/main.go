package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		delResp *clientv3.DeleteResponse
		kvpair  *mvccpb.KeyValue
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

	// withPrevKV 需要带上 不然 delResp.PrevKvs 为空
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job3", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	} else {
		if len(delResp.PrevKvs) != 0 {
			for _, kvpair = range delResp.PrevKvs {
				fmt.Println("del", string(kvpair.Key), string(kvpair.Value))
			}
		} else {
			fmt.Println("del nothing")
		}
	}

}
