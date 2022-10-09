package main

import (
	"fmt"
	"time"

	"github.com/ronething/learngo/retriever/mock"
	"github.com/ronething/learngo/retriever/real"
)

type Retiever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://blog.ronething.cn"

func download(r Retiever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ronething",
			"github": "ronething",
		})
}

// 接口的组合
type RetriverPoster interface {
	Retiever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked ronething.com",
	})
	return s.Get(url)
}

func main() {
	var r Retiever

	retriver := mock.Retriever{"this is a fack ronething.com"}
	r = &retriver
	inspect(r)

	// Type assertion
	mockRetriever := r.(*mock.Retriever)
	fmt.Println(mockRetriever.Contents)
	fmt.Println(session(&retriver))

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	//fmt.Println(download(r))
}

func inspect(r Retiever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
