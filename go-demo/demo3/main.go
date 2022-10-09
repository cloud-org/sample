package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

// 代表一个任务
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	//	 执行多个任务
	//	启动调度协程定时检查所有 cron 任务 谁过期就执行谁
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob // key: task name
	)

	scheduleTable = make(map[string]*CronJob)

	now = time.Now()
	//1、定义 2 个 CronJob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表里
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/10 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表里
	scheduleTable["job2"] = cronJob

	// 2、启动调度协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				//	判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					//	启动一个协程 执行这个任务
					go func(jobName string) {
						fmt.Println("exec ", jobName)
					}(jobName)

					//	计算下一次调度执行时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "next time is", cronJob.nextTime)
				}

			}

			// 睡眠 100ms
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:
			}

			//time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(100 * time.Second)
}
