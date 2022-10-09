package main

import (
	"fmt"

	"github.com/ronething/learngo/queue"
)

func main() {

	q := queue.Quque{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
