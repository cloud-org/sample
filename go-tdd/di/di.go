// author: ashing
// time: 2020/4/17 12:24 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//func main() {
//	Greet(os.Stdout, "Elodie")
//}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe("127.0.0.1:5000", http.HandlerFunc(MyGreeterHandler))
}
