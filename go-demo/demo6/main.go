package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		kv             clientv3.KV
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse // 只读 chan

		err error
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second, // 连接超时时间
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//申请一个租约
	lease = clientv3.NewLease(client)

	//	申请一个 10s 的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//put kv 与租约关联
	leaseId = leaseGrantResp.ID

	//自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}
	//处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效")
					goto END
				} else {
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}
	END:
	}()

	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("put succeed", putResp.Header.Revision)

	//	通过 watch 观察变化

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv 过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
