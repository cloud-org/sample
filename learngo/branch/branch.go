package main

import (
	"fmt"
	"io/ioutil"
)

// switch 不需要 break
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
		//default:
		//	panic(fmt.Sprintf("Wrong score: %d", score))
		//
	}
	return g
}

func main() {
	const filename = "hello.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	/*
		if 条件里面可以赋值 变量的作用域就在这个 if 语句里
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(10),
		grade(99),
		//grade(-1),
		//grade(101),
	)
}
