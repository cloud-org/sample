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

package svc

import (
	"log"
	"runtime/debug"
	"strings"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Agent struct {
	Id       AgentId
	RegionId RegionId
	MsgChan  chan []byte
}

func NewSubAgent(id string, regionId string) *Agent {
	return &Agent{
		Id:       AgentId(id),
		RegionId: RegionId(regionId),
		MsgChan:  make(chan []byte, 10),
	}
}

type AgentId string
type RegionId string
type Agents map[AgentId]*Agent

type RegionBroad struct {
	prefix       string // etcd watch key prefix
	RegionAgents map[RegionId]Agents
	addC         chan *Agent
	removeC      chan *Agent
	watchC       clientv3.WatchChan
	stopC        chan struct{}
}

//trimPrefix 过滤出对应变动的 regionId
func (r *RegionBroad) trimPrefix(key string) RegionId {
	return RegionId(strings.TrimPrefix(key, r.prefix))
}

func NewRegionBroad(key string, watchC clientv3.WatchChan) *RegionBroad {
	return &RegionBroad{
		prefix:       key,
		RegionAgents: make(map[RegionId]Agents),
		addC:         make(chan *Agent),
		removeC:      make(chan *Agent),
		watchC:       watchC,
		stopC:        make(chan struct{}),
	}
}

func (r *RegionBroad) Loop() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("广播器 loop 发生异常, err: %v\n", err)
			debug.PrintStack()
		}
	}()

	for {
		select {
		case watchResp := <-r.watchC:
			for _, event := range watchResp.Events {
				log.Printf("key: %v value: %v\n", string(event.Kv.Key), string(event.Kv.Value))
				switch event.Type {
				case mvccpb.PUT:
					regionId := r.trimPrefix(string(event.Kv.Key))
					// 将配置进行广播
					agents := r.RegionAgents[regionId]
					log.Printf("agents count %v\n", len(agents))
					for _, agent := range agents {
						agent.MsgChan <- event.Kv.Value
					}
				case mvccpb.DELETE:
					// TODO: 删除事件 暂时不需要用到
					log.Printf("删除事件, key: %v value: %v\n", string(event.Kv.Key), string(event.Kv.Value))
				}
			}
		case agent := <-r.addC:
			agents, ok := r.RegionAgents[agent.RegionId]
			if !ok {
				r.RegionAgents[agent.RegionId] = make(map[AgentId]*Agent)
				agents = r.RegionAgents[agent.RegionId]
			}
			agents[agent.Id] = agent
		case agent := <-r.removeC:
			agents, ok := r.RegionAgents[agent.RegionId]
			if ok {
				delete(agents, agent.Id)
			}
		case <-r.stopC:
			log.Println("广播器退出")
			return
		}
	}
}

func (r *RegionBroad) AddAgent(agent *Agent) {
	r.addC <- agent
}

func (r *RegionBroad) RemoveAgent(agent *Agent) {
	r.removeC <- agent
}

func (r *RegionBroad) Stop() {
	r.stopC <- struct{}{}
}
