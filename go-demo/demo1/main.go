package main

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

type result struct {
	output []byte
	err    error
}

func main() {

	// cancelFunc 无法正常取消 通过 group pid 取消
	//	exec one cmd for 2s, but kill it after 1s
	var (
		ctx context.Context
		//cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
		err        error
	)

	resultChan = make(chan *result, 1000)

	ctx, _ = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 60;echo hello")
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		output, err = cmd.CombinedOutput()

		resultChan <- &result{
			output: output,
			err:    err,
		}

	}()

	time.Sleep(1 * time.Second)

	//cancelFunc()

	if err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
		fmt.Println(err)
		return
	}

	res = <-resultChan

	fmt.Printf("output is %s\n", string(res.output))

	fmt.Println(res.err)

}
