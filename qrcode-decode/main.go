package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tuotoo/qrcode"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `qrcode-decode tool
Usage: qrcode [-h help]
Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.StringVar(&filePath, "p", "./qrcode.png", "图片文件路径")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		usage()
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("打开文件 %s 失败, %v\n", filePath, err.Error())
		return
	}
	defer file.Close()
	qrMatrix, err := qrcode.Decode(file)
	if err != nil {
		log.Println("识别二维码失败", err.Error())
		return
	}
	log.Printf("识别链接为 %s\n", qrMatrix.Content)
}
