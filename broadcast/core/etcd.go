/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package core

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"google.golang.org/grpc"
)

type EtcdMgr struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	watcher clientv3.Watcher
}

func InitEtcdMgr(endpoints []string, dialTimeout time.Duration) (jobMgr *EtcdMgr) {

	//	初始化配置
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
		// 需要加上 DialOptions 不然不会有 err
		// more: https://github.com/etcd-io/etcd/issues/9877
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}

	client, err := clientv3.New(config)
	if err != nil {
		log.Fatalf("初始化 etcd 配置失败, err: %v\n", err)
		return
	}

	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	watcher := clientv3.NewWatcher(client)

	return &EtcdMgr{
		client:  client,
		kv:      kv,
		lease:   lease,
		watcher: watcher,
	}
}

func (e *EtcdMgr) Close() {
	_ = e.client.Close()
	_ = e.lease.Close()
	_ = e.watcher.Close()
}

//CreateWatchChan 创建 watch 通道
func (e *EtcdMgr) CreateWatchChan(ctx context.Context, key string) (clientv3.WatchChan, error) {
	// constant.ChannelConfigDir
	getResp, err := e.kv.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	watchRevision := getResp.Header.Revision + 1
	watchChan := e.watcher.Watch(ctx, key, clientv3.WithRev(watchRevision), clientv3.WithPrefix())
	return watchChan, nil
}

func (e *EtcdMgr) Put(ctx context.Context, key string, value string) error {
	putResp, err := e.kv.Put(ctx, key, value, clientv3.WithPrevKV())
	if err != nil {
		return err
	}
	if putResp.PrevKv != nil {
		log.Printf("prev is: %v\n", string(putResp.PrevKv.Value))
	}
	return nil
}
