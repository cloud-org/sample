package main

import (
	"log"
	"multi-timezone-scheduler/scheduler"
	"os"
	"os/signal"
	"syscall"
)

func initTask() []scheduler.Task {
	tasks := make([]scheduler.Task, 0)
	// func NewTask(tid string, name string, expression string, timezone string, command string) *Task
	task1 := scheduler.NewTask(
		"task-1", "task-1", "* * * * *", "Asia/Shanghai", "echo hello shanghai")
	tasks = append(tasks, *task1)
	task2 := scheduler.NewTask(
		"task-2", "task-2", "* * * * *", "Asia/Tokyo", "echo hello tokyo")
	tasks = append(tasks, *task2)

	return tasks
}

func main() {
	_ = scheduler.InitScheduler(initTask())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// DONE: 优雅关停
	for {
		s := <-c
		log.Printf("[main] 捕获信号 %s", s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// 停止调度器 并等待正在 running 的任务执行结束
			scheduler.StopScheduler()
			return // 很重要 不然程序无法退出
		case syscall.SIGHUP:
			log.Printf("[main] 终端断开信号，忽略")
		default:
			log.Printf("[main] 其他信号")
		}
	}

}
