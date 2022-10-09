package core

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

var myerr = errors.New("custom err")

func TestNewTimeWheel(t *testing.T) {
	tw, err := NewTimeWheel(1*time.Second, 3)
	if err != nil {
		panic(err)
	}

	tw.Start() // 开始调度
	defer tw.Stop()

	// 先声明
	var task1 *Task

	wg := sync.WaitGroup{}

	wg.Add(2)
	count := 1 // 跑两次
	task1 = tw.AddCron(5*time.Second, func() {
		defer wg.Done()
		// 业务逻辑 查询出相应的逻辑
		log.Println("业务....")
		if err := RunDelayTask(task1, count); err != nil {
			if err == myerr {
				t.Log("删除自己")
				tw.Remove(task1)
			}
		}
		count--
	})

	wg.Wait()

	fmt.Println("结束")
}

func RunDelayTask(task1 *Task, count int) error {
	fmt.Println(task1)
	if count == 0 {
		return myerr
	}
	return nil
}

func TestTimeWheelRemoveTask(t *testing.T) {
	tw, err := NewTimeWheel(100*time.Millisecond, 10000)
	if err != nil {
		fmt.Println(err)
	}
	tw.Start()
	defer tw.Stop()
	task1 := tw.Add(3*time.Second, func() {
		fmt.Println("xxx")
	})
	time.Sleep(1 * time.Second) // 等一秒再 remove 确保是 add 事件先被通道获取
	tw.Remove(task1)
	var wg sync.WaitGroup
	wg.Add(1)
	tw.Add(3*time.Second, func() {
		defer wg.Done()
		fmt.Println("xxx2")
	})
	wg.Wait()
}

func TestTimeWheelCalc(t *testing.T) {
	tw, err := NewTimeWheel(1*time.Second, 3)
	if err != nil {
		fmt.Println(err)
	}
	tw.Start()
	defer tw.Stop()
	var wg sync.WaitGroup
	wg.Add(10)
	tw.AddCron(3*time.Second, func() {
		defer wg.Done()
		log.Println("hello world")
	})
	wg.Wait()
}

func TestTimeWheelAdd(t *testing.T) {
	tw, err := NewTimeWheel(1*time.Second, 3)
	if err != nil {
		fmt.Println(err)
	}
	tw.Start()
	defer tw.Stop()
	var wg sync.WaitGroup
	wg.Add(1)
	tw.Add(0*time.Second, func() {
		defer wg.Done()
		log.Println("hello world")
	})
	wg.Wait()
}
