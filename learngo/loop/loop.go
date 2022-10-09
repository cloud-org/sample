package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	pwd, _ := os.Getwd()
	pwd += "/loop/"
	file, err := os.Open(pwd + filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

/*
	相当于 while True
	go 没有 while
*/
func forever() {
	for {
		fmt.Println("hello world")
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(128),
	)
	printFile("hello.txt")

	s := `很秀
	这个字符串可以
	跨行
	`
	printFileContents(strings.NewReader(s))

	//forever()
}
