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

package main

import (
	"broadcast/core"
	"broadcast/svc"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var regionBroad *svc.RegionBroad

func main() {
	// init etcdMgr
	etcdMgr := core.InitEtcdMgr([]string{"127.0.0.1:2379"}, 5*time.Second)

	watchPrefix := "/config/"
	// init broad mgr
	watchC, err := etcdMgr.CreateWatchChan(context.TODO(), watchPrefix)
	if err != nil {
		log.Println(err)
		return
	}
	regionBroad = svc.NewRegionBroad(watchPrefix, watchC)
	go regionBroad.Loop()
	defer regionBroad.Stop()

	// 广播器加入 两个 agent 属于相同 region
	go addAgent("agent_1", "region_123456")
	go addAgent("agent_2", "region_123456")

	stopC := make(chan os.Signal)
	signal.Notify(stopC, syscall.SIGINT, syscall.SIGTERM)
	<-stopC
}

func addAgent(xHost, regionId string) {
	log.Println("add agent", xHost, regionId)
	subAgent := svc.NewSubAgent(xHost, regionId)
	regionBroad.AddAgent(subAgent)
	for {
		select {
		case msg := <-subAgent.MsgChan:
			log.Printf("%s 接收到新的配置，%v\n", xHost, string(msg))
		}
	}
}
