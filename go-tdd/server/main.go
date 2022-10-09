// author: ashing
// time: 2020/4/21 12:56 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
