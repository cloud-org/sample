package scheduler

import (
	"context"
	"delay-task-scheduler/core"
	"log"
	"sync"
	"time"
)

var delayTaskScheduler *DelayTaskScheduler // 可作为全局延时任务调度实例

type DelayTaskScheduler struct {
	tw      *core.TimeWheel       // 时间轮实例
	tasks   map[string]*core.Task // 维护 delayTask id 和时间轮中对应 task 的映射关系
	addC    chan *DelayTask       // 添加任务事件通道
	removeC chan *DelayTask       // 删除任务事件通道
	wg      sync.WaitGroup        // 执行任务 waitGroup 用于停止调度器的时候使用
	stopC   chan struct{}         // 接收暂停事件通道
}

func NewDelayTaskScheduler(tw *core.TimeWheel) *DelayTaskScheduler {
	return &DelayTaskScheduler{
		tw:      tw,
		tasks:   make(map[string]*core.Task),
		addC:    make(chan *DelayTask, 10),
		removeC: make(chan *DelayTask, 10),
		wg:      sync.WaitGroup{},
		stopC:   make(chan struct{}),
	}
}

//loop 开启循环处理各种事件
func (d *DelayTaskScheduler) loop() {
	timer := time.NewTimer(1 * time.Minute)
	for {
		select {
		case task := <-d.addC:
			d.add(task)
		case task := <-d.removeC:
			d.remove(task)
		case <-d.stopC:
			d.tw.Stop()
			log.Printf("延时任务调度器退出")
			return
		case <-timer.C:
			log.Printf("剩余的延时任务数量: %v", len(d.tasks))
			timer.Reset(1 * time.Minute)
		}
	}
}

func StopDelayTaskScheduler() *sync.WaitGroup {
	delayTaskScheduler.stopC <- struct{}{}
	return &delayTaskScheduler.wg
}

func InitDelayTaskScheduler(ctx context.Context) error {
	tw, err := core.NewTimeWheel(1*time.Second, 60) // 支持分钟级别
	if err != nil {
		return err
	}
	delayTaskScheduler = NewDelayTaskScheduler(tw)
	delayTaskScheduler.tw.Start() // 启动时间轮
	go delayTaskScheduler.loop()  // 启动循环

	// 延时任务初始化
	if err = initDelayTask(ctx); err != nil {
		log.Printf("初始化延时任务失败")
		return err
	}

	return nil
}

func initDelayTask(ctx context.Context) error {
	delayTasks := make([]DelayTask, 0)

	// TODO: 实际可从数据库中获取延时任务

	for i := 0; i < len(delayTasks); i++ {
		task := delayTasks[i]
		log.Printf(
			"延时任务 %s-%s 信息 disable: %v, state: %v",
			task.Tid,
			task.Name,
			task.Disable,
			task.State,
		)
		if !task.Disable && task.State == Pending { // 需要添加
			// calc delay
			delay := task.Delay - (time.Now().Unix() - task.UpdateAt)
			if delay <= 0 {
				delay = 0
			}
			task.Delay = delay
			delayTaskScheduler.addTask(&task)
		} else {
			log.Printf("延时任务无需添加")
		}
	}

	return nil
}

func (d *DelayTaskScheduler) CreateDelayTask(task *DelayTask) {
	if task.Disable {
		log.Printf("[scheduler] 延时任务 %s-%s 不需要添加到延时任务调度器中\n", task.Tid, task.Name)
	} else {
		log.Printf("[scheduler] 延时任务 %s-%s 添加到延时任务调度器中\n", task.Tid, task.Name)
		d.addTask(task)
	}
	return
}

func (d *DelayTaskScheduler) DeleteDelayTask(task *DelayTask) {
	d.removeTask(task)
	return
}

func (d *DelayTaskScheduler) DisableDelayTask(oldTask *DelayTask, disable bool) {
	if oldTask.Disable == disable {
		log.Printf("状态没有变动，无需修改")
		return
	}
	if oldTask.Disable && !disable && oldTask.State == Pending { // true -> false
		d.addTask(oldTask)
		return
	}
	if !oldTask.Disable && disable && oldTask.State == Pending { // false -> true
		d.removeTask(oldTask)
		return
	}
}

func (d *DelayTaskScheduler) addTask(task *DelayTask) {
	d.addC <- task
	return
}

func (d *DelayTaskScheduler) removeTask(task *DelayTask) {
	d.removeC <- task
	return
}

func (d *DelayTaskScheduler) add(task *DelayTask) {
	var timeWheelTask *core.Task
	timeWheelTask = d.tw.Add(time.Duration(task.Delay)*time.Second, func() {
		d.wg.Add(1)
		defer d.wg.Done()
		// TODO: 业务逻辑
		log.Println("假装正在执行任务...")
		time.Sleep(1 * time.Second)
		log.Println("假装执行任务完成...")

		// 执行成功之后，将调度器中的任务删除
		delayTaskScheduler.removeTask(task)
	})

	d.tasks[task.Tid] = timeWheelTask
	log.Printf("添加延时任务 %s-%s 成功...", task.Tid, task.Name)

	return
}

func (d *DelayTaskScheduler) remove(task *DelayTask) {
	log.Printf("移除延时任务 %s-%s\n", task.Tid, task.Name)
	twTask, ok := d.tasks[task.Tid]
	if !ok {
		log.Printf("获取不到 task，无需删除")
		return
	}
	log.Printf("时间轮 task is %v\n", twTask)
	delete(d.tasks, task.Tid) // 没有的 key 也不会报错
	d.tw.Remove(twTask)

	return
}
