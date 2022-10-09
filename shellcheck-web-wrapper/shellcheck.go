package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Json1 结果结构体
// shellcheck --format=json1
type Json1 struct {
	Comments []Comments `json:"comments"`
}

//Comments 具体的错误
type Comments struct {
	//File      string      `json:"file"` // 文件名注释
	//Fix       interface{} `json:"fix"` // 修复建议注释
	Line      int    `json:"line"`
	EndLine   int    `json:"endLine"`
	Column    int    `json:"column"`
	EndColumn int    `json:"endColumn"`
	Level     string `json:"level"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

//ShellCheck 传入脚本字符串，通过 sc 检测
func ShellCheck(script string) (*Json1, error) {

	// 1、写入临时文件
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sc*.sh") // * 会被替换为随机字符串
	if err != nil {
		log.Printf("create tempFile err: %v\n", err)
		return nil, err
	}
	log.Printf("tmp file name is %v\n", tmpFile.Name())
	_, err = tmpFile.Write([]byte(script))
	if err != nil {
		log.Printf("write err: %v\n", err)
		return nil, err
	}
	defer func() {
		_ = os.Remove(tmpFile.Name())
	}()

	var stdOutBuf bytes.Buffer
	var stdErrBuf bytes.Buffer
	var json1 Json1
	// 2、执行 shellcheck 并进行反序列化
	c := exec.Command("/usr/bin/shellcheck", tmpFile.Name(), "--format=json1")
	c.Stdout = &stdOutBuf
	c.Stderr = &stdErrBuf

	err = c.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			//log.Printf("err is %+v\n", exitError)
			//log.Printf("err is %+v\n", exitError.ExitCode())
			if exitError.ExitCode() < 128 && stdOutBuf.String() != "" { // 正常
				//log.Printf("stdout is %v\n", stdOutBuf.String())
				err = json.Unmarshal(stdOutBuf.Bytes(), &json1)
				if err != nil {
					return nil, err
				}
				return &json1, nil
			}
		}
		// 写入错误信息
		stdErrBuf.WriteString(err.Error())
		log.Printf("stderr is %v\n", stdErrBuf.String())
		return nil, err
	}

	// 正常来讲不会走到这里
	if stdOutBuf.String() != "" {
		err = json.Unmarshal(stdOutBuf.Bytes(), &json1)
		if err != nil {
			return nil, err
		}

		return &json1, nil
	}

	return nil, errors.New("未知错误")

}
