package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)
	mux = http.NewServeMux()
	mux.HandleFunc("/post", handlePost)

	addr := "127.0.0.1:8080"

	listener, _ = net.Listen("tcp", addr)

	httpServer = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}

	go httpServer.Serve(listener)

	fmt.Printf("start server at addr: %s\n", addr)

	for {
		time.Sleep(1 * time.Second)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf(r.Header.Get("Content-Type"))
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, "{\"msg\":\"success\"}")
	fmt.Printf("%s %s\n", info, body)
}
