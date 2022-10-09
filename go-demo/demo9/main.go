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
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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

	//	分布式乐观锁
	//op 操作
	//txn 事务: if else then

	// 1、上锁(创建租约，自动续租，拿着租约去抢占一个 key)

	//申请一个租约
	lease = clientv3.NewLease(client)

	//	申请一个 10s 的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//put kv 与租约关联
	leaseId = leaseGrantResp.ID

	ctx, cancelFunc = context.WithCancel(context.TODO())

	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	//自动续租
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}
	END:
	}()

	//	if then else

	kv = clientv3.NewKV(client)

	//创建一个事务
	txn = kv.Txn(context.TODO())

	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "xxx", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9"))

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	//	判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//2、处理业务
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)

	//3、释放锁(取消自动续租，立即释放租约)
	// defer

}
