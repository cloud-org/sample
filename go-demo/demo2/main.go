package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)
	// cron expr: sec min hour day month week year
	// 粒度更细

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	now = time.Now()

	// 0 5 10 15...
	nextTime = expr.Next(now)

	fmt.Println(now)
	fmt.Println(nextTime)

	// 等到计时器超时 就可以执行下一个任务了
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("exec")
	})

	// 阻塞一下
	time.Sleep(5 * time.Second)

}
