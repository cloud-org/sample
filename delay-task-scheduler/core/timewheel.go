package core

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

const (
	modeIsCircle  = true
	modeNotCircle = false

	modeIsAsync  = true
	modeNotAsync = false
)

type taskID int64

type Task struct {
	delay    time.Duration
	id       taskID
	round    int
	callback func()

	async  bool
	stop   bool
	circle bool
	first  bool // 是否是首次添加
}

// Reset for sync.Pool
func (t *Task) Reset() {
	t.round = 0
	t.callback = nil

	t.async = false
	t.stop = false
	t.circle = false
	t.first = true
}

type optionCall func(*TimeWheel) error

func SetSyncPool(state bool) optionCall {
	return func(o *TimeWheel) error {
		o.syncPool = state
		return nil
	}
}

type TimeWheel struct {
	randomID int64

	tick   time.Duration
	ticker *time.Ticker

	bucketsNum    int
	buckets       []map[taskID]*Task // key: added item, value: *Task
	bucketIndexes map[taskID]int     // key: added item, value: bucket position

	currentIndex int

	onceStart sync.Once

	addC    chan *Task
	removeC chan *Task
	stopC   chan struct{}

	exited   bool
	syncPool bool
}

// NewTimeWheel create new time wheel
func NewTimeWheel(tick time.Duration, bucketsNum int, options ...optionCall) (*TimeWheel, error) {
	if tick.Seconds() < 0.1 {
		return nil, errors.New("invalid params, must tick >= 100 ms")
	}
	if bucketsNum <= 0 {
		return nil, errors.New("invalid params, must bucketsNum > 0")
	}

	tw := &TimeWheel{
		// tick
		tick: tick,

		// store
		bucketsNum:    bucketsNum,
		bucketIndexes: make(map[taskID]int, 1024*100),
		buckets:       make([]map[taskID]*Task, bucketsNum),
		currentIndex:  0,

		// signal
		addC:    make(chan *Task, 1024*5),
		removeC: make(chan *Task, 1024*2),
		stopC:   make(chan struct{}),
	}

	for i := 0; i < bucketsNum; i++ {
		tw.buckets[i] = make(map[taskID]*Task, 16)
	}

	for _, op := range options {
		op(tw)
	}

	return tw, nil
}

// Start start the time wheel
func (tw *TimeWheel) Start() {
	tw.onceStart.Do(
		func() {
			tw.ticker = time.NewTicker(tw.tick)
			go tw.scheduler() // 启动调度协程
		},
	)
}

// 调度协程
func (tw *TimeWheel) scheduler() {
	queue := tw.ticker.C

	for {
		select {
		case <-queue:
			log.Println("handle")
			tw.handleTick()
		case task := <-tw.addC:
			log.Println("add task", task)
			tw.put(task)
		case task := <-tw.removeC:
			//log.Println("remove task", task)
			tw.remove(task)
		case <-tw.stopC:
			tw.exited = true
			tw.ticker.Stop()
			return
		}
	}
}

// Stop stop the time wheel
func (tw *TimeWheel) Stop() {
	tw.stopC <- struct{}{}
}

func (tw *TimeWheel) collectTask(task *Task) {
	index := tw.bucketIndexes[task.id]
	delete(tw.bucketIndexes, task.id)
	delete(tw.buckets[index], task.id)

	if tw.syncPool {
		defaultTaskPool.put(task)
	}
}

func (tw *TimeWheel) handleTick() {
	bucket := tw.buckets[tw.currentIndex]
	for k, task := range bucket {
		if task.stop {
			tw.collectTask(task)
			continue
		}

		if bucket[k].round > 0 {
			bucket[k].round--
			continue
		}

		// 如果不是异步就会等
		if task.async {
			go task.callback()
		} else {
			// optimize gopool
			task.callback()
		}

		// circle
		if task.circle == true {
			tw.collectTask(task)
			tw.put(task)
			continue
		}

		// gc
		tw.collectTask(task)
	}

	if tw.currentIndex == tw.bucketsNum-1 {
		tw.currentIndex = 0
		return
	}

	tw.currentIndex++
}

// Add add an task
func (tw *TimeWheel) Add(delay time.Duration, callback func()) *Task {
	return tw.addAny(delay, callback, modeNotCircle, modeIsAsync) // 默认异步
}

// AddCron add interval task
func (tw *TimeWheel) AddCron(delay time.Duration, callback func()) *Task {
	return tw.addAny(delay, callback, modeIsCircle, modeIsAsync)
}

func (tw *TimeWheel) addAny(delay time.Duration, callback func(), circle, async bool) *Task {
	// TODO: delay < tw.tick 是否要允许
	if delay < tw.tick {
		log.Println("任务延迟时间小于最小刻度，设置为最小刻度")
		delay = tw.tick
	}

	id := tw.genUniqueID()

	var task *Task
	if tw.syncPool {
		task = defaultTaskPool.get()
	} else {
		task = new(Task)
	}

	task.delay = delay
	task.id = id
	task.callback = callback
	task.circle = circle
	task.async = async // refer to src/runtime/time.go
	task.first = true  // 首次添加

	tw.addC <- task
	return task
}

func (tw *TimeWheel) put(task *Task) {
	tw.store(task)
}

func (tw *TimeWheel) store(task *Task) {
	round := tw.calculateRound(task)
	index := tw.calculateIndex(task)
	if task.first { // 首次添加完之后要设置 first 为 false，如果是 cron 的话 之后会有不同的计算方法
		task.first = false
	}

	task.round = round

	tw.bucketIndexes[task.id] = index
	tw.buckets[index][task.id] = task
}

func (tw *TimeWheel) calculateRound(task *Task) (round int) {
	delay := task.delay
	// 减掉一个 tick 计算出来的 round 才是准确的
	// 例如 0-9 然后设置 delay 为 10s，其实 round 应该不是为 1， 而是为 0，然后 index 为 9
	delaySeconds := (delay - tw.tick).Seconds()
	tickSeconds := tw.tick.Seconds()
	round = int(delaySeconds / tickSeconds / float64(tw.bucketsNum))
	log.Println("round", round)
	return
}

func (tw *TimeWheel) calculateIndex(task *Task) (index int) {
	delay := task.delay
	delaySeconds := (delay - tw.tick).Seconds() // 减掉一个 tick 计算出来的 index 才是准确的
	tickSeconds := tw.tick.Seconds()
	log.Println(delaySeconds, tickSeconds, tw.currentIndex, tw.bucketsNum)
	if task.first {
		index = (int(float64(tw.currentIndex) + delaySeconds/tickSeconds)) % tw.bucketsNum
	} else { // cron
		currentIndex := (tw.currentIndex + 1) % tw.bucketsNum
		index = (int(float64(currentIndex) + delaySeconds/tickSeconds)) % tw.bucketsNum
	}
	log.Println("index", index)
	return
}

func (tw *TimeWheel) Remove(task *Task) {
	tw.removeC <- task
	return
}

func (tw *TimeWheel) remove(task *Task) {
	tw.collectTask(task)
}

func (tw *TimeWheel) genUniqueID() taskID {
	id := atomic.AddInt64(&tw.randomID, 1)
	return taskID(id)
}
