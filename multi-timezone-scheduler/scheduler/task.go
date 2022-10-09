package scheduler

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/pkg/errors"
)

// 当前任务
type Task struct {
	Tid        string `json:"tid" bson:"tid"`               // 任务 id
	Name       string `json:"name" bson:"name"`             // task 名字
	Expression string `json:"expression" bson:"expression"` // 表达式 支持@every [1s | 1m | 1h ] 参考 cron
	Timezone   string `json:"timezone" bson:"timezone"`     // 新增时区配置
	Command    string `json:"command"`                      // 执行命令
}

func NewTask(tid string, name string, expression string, timezone string, command string) *Task {
	return &Task{Tid: tid, Name: name, Expression: expression, Timezone: timezone, Command: command}
}

func (t *Task) Run() error {
	var (
		stdOutBuf bytes.Buffer
		stdErrBuf bytes.Buffer
	)

	if t.Command == "" {
		return errors.New("please do not input the empty command")
	}

	log.Printf("now will run the task [%s-%s]", t.Tid, t.Name)

	c := exec.Command("/bin/bash", "-c", t.Command)
	c.Stdout = &stdOutBuf
	c.Stderr = &stdErrBuf

	e := c.Run()
	if e != nil {
		log.Printf("run task err: %v", e)
		// 写入错误信息
		stdErrBuf.WriteString(e.Error())
		return e
	}

	log.Printf("stdout: %v\n", stdOutBuf.String())
	log.Printf("stderr: %v\n", stdErrBuf.String())

	return nil

}
