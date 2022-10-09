// author: ashing
// time: 2020/7/5 11:18 下午
// mail: axingfly@gmail.com
// Less is more.
// context 通知退出策略

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")
	ctxb, _ := context.WithTimeout(ctxa, 3*time.Second) // 三秒过期
	go work(ctxb, "work2")
	ctxc := context.WithValue(ctxb, "name", "panda") // 三秒过期
	go workWithValue(ctxc, "work3")
	time.Sleep(5 * time.Second)
	//cancel3() // 会不会传播到整棵树 answer: 不会
	cancel()                    // 手动调用 cancel
	time.Sleep(3 * time.Second) // 等待信息打印
}

// work 监听 context
func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running\n", name)
			time.Sleep(1 * time.Second)
		}

	}
}

// workWithValue context with Value
func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("name").(string)
			fmt.Printf("%s is running value=%s\n", name, value)
			time.Sleep(1 * time.Second)
		}

	}
}
