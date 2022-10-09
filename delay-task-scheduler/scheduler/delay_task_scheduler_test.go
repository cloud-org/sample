package scheduler

import (
	"context"
	"testing"
	"time"
)

func TestInitDelayTaskScheduler(t *testing.T) {
	err := InitDelayTaskScheduler(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}

	delayTaskScheduler.CreateDelayTask(&DelayTask{
		Tid:   "123456",
		Name:  "这是一个延时 10s 的任务",
		Delay: 10,
	})

	time.Sleep(12 * time.Second) // 手动延时足够的时间 实际使用的话一般会是嵌入某个服务中常驻

	wg := StopDelayTaskScheduler()
	wg.Wait() // 如果有正在执行的任务 会等待任务执行完成之后才退出
}
