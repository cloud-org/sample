package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 今天是个好日子" // utf-8
	fmt.Println(s)
	//fmt.Println(len(s)) // len 获得字节长度
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune conut: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Print(ch, size, " ")
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
		fmt.Println()
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
