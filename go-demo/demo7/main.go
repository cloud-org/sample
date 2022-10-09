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
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchChan          <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
		ctx                context.Context
		cancelFunc         context.CancelFunc
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

	// 模拟 etcd 中 kv 的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")

			kv.Delete(context.TODO(), "/cron/jobs/job7")

			time.Sleep(1 * time.Second)
		}
	}()

	//	 先 get 当前的值
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}

	//key exist
	if len(getResp.Kvs) != 0 {
		fmt.Println("current is ", string(getResp.Kvs[0].Value))
	}

	//	current revision id monotonic increase
	watchStartRevision = getResp.Header.Revision + 1

	// create a watcher
	watcher = clientv3.NewWatcher(client)

	fmt.Println("watch from revision", watchStartRevision)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	watchChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	for watchResp = range watchChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}
