package scheduler

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/robfig/cron/v3"
)

var cronScheduler *CronScheduler

type CronScheduler struct {
	mu         sync.Mutex
	schedulers map[string]*cron.Cron
	tasks      map[string]cron.EntryID // worker 中的定时任务 map
}

//NewCronScheduler 构造函数
func NewCronScheduler() *CronScheduler {
	return &CronScheduler{
		mu:         sync.Mutex{},
		schedulers: make(map[string]*cron.Cron),
		tasks:      make(map[string]cron.EntryID),
	}
}

//GetTimezoneScheduler 获取时区定时器
func (c *CronScheduler) GetTimezoneScheduler(timezone string) *cron.Cron {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.schedulers[timezone] // nil or scheduler
}

//AddTimezoneScheduler
func (c *CronScheduler) AddTimezoneScheduler(timezone string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if scheduler, ok := c.schedulers[timezone]; ok {
		log.Printf("该时区 %s 定时器已创建\n", timezone)
		scheduler.Start() // 重复 start 也没事 幂等
		return nil        // 直接返回
	}
	// 创建对应的时区定时器
	optLogs := cron.WithLogger(
		cron.VerbosePrintfLogger(
			log.New(os.Stdout, "[Cron]: ", log.LstdFlags)))

	optParser := cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	))

	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("[scheduler] load time location err: %v\n", err)
		return err
	}
	optLocation := cron.WithLocation(location)
	scheduler := cron.New(optLogs, optParser, optLocation)
	scheduler.Start() // 重复 start 也没事 幂等
	c.schedulers[timezone] = scheduler
	log.Printf("创建时区 %s 定时器成功\n", timezone)
	return nil
}

//RemoveTimezoneScheduler
func (c *CronScheduler) RemoveTimezoneScheduler(timezone string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.schedulers, timezone)
}

//StartTimezoneScheduler
func (c *CronScheduler) StartTimezoneScheduler(timezone string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	scheduler, ok := c.schedulers[timezone]
	if !ok {
		return errors.New(fmt.Sprintf("该时区 %s 定时器未创建", timezone))
	}
	scheduler.Start()
	return nil
}

//StartAllScheduler
func (c *CronScheduler) StartAllScheduler() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, v := range c.schedulers {
		v.Start()
	}
}

func (c *CronScheduler) StopAllScheduler() {
	c.mu.Lock()
	defer c.mu.Unlock()
	wg := sync.WaitGroup{}
	for key, value := range c.schedulers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, key string, value *cron.Cron) {
			defer wg.Done()
			ctx := value.Stop()
			for {
				select {
				case <-ctx.Done():
					goto End
				default:
					time.Sleep(1 * time.Second)
					log.Printf("%s 定时器仍在退出", key)
				}
			}
		End:
			log.Printf("%s 定时器停止完成", key)
		}(&wg, key, value)
	}
	wg.Wait()
}

func (c *CronScheduler) GetTaskEntryId(tid string) cron.EntryID {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.tasks[tid] // if key not exists, value is zero
}

func (c *CronScheduler) PutTaskEntryId(tid string, id cron.EntryID) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tasks[tid] = id
}

//RemoveTask 删除定时器中的定时任务
func (c *CronScheduler) RemoveTask(t *Task) {
	scheduler := c.GetTimezoneScheduler(t.Timezone)
	if scheduler == nil {
		log.Printf("获取不到对应时区 %s 定时器\n", t.Timezone)
		return
	}
	// 获取任务 cron.EntryId
	entryId := c.GetTaskEntryId(t.Tid)
	log.Printf("task %s entryId is %v\n", t.Tid, entryId)
	if entryId != 0 {
		scheduler.Remove(entryId)
	}
}

func (c *CronScheduler) AddTask(t *Task) error {
	f := func() {
		if e := t.Run(); e != nil {
			log.Printf("[scheduler] exec task %s err: %v", t.Tid, e)
		}
	}

	scheduler := c.GetTimezoneScheduler(t.Timezone)
	if scheduler == nil {
		err := c.AddTimezoneScheduler(t.Timezone) // 新建一个
		if err != nil {
			return err
		}
	}
	scheduler = c.GetTimezoneScheduler(t.Timezone) // 这次理论上肯定能获取到
	entryId, err := scheduler.AddFunc(t.Expression, f)
	if err != nil {
		log.Printf("[scheduler] add func to scheduler err: %v", err)
		return err
	}

	c.PutTaskEntryId(t.Tid, entryId)
	log.Printf("[scheduler] add the job of %s\n", t.Name)

	return nil
}

//InitScheduler 初始化定时器以及任务
func InitScheduler(tasks []Task) error {
	// 初始化 cronScheduler
	cronScheduler = NewCronScheduler()
	// 将任务加入时区定时器
	// 兼容一下原先的数据 GMT+8
	for i := 0; i < len(tasks); i++ {
		// 默认清空之前的状态
		t := tasks[i]
		if err := cronScheduler.AddTask(&t); err != nil {
			log.Fatalf("[scheduler] init the task err:%v", err)
		}
	}
	return nil
}

//StopScheduler 停止调度器
func StopScheduler() {
	cronScheduler.StopAllScheduler()
}
