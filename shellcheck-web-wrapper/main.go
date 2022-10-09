package main

import "log"

func main() {
	echoServer, err := CreateEngine()
	if err != nil {
		log.Println(err)
		return
	}

	if err = echoServer.Start(":7777"); err != nil {
		log.Println(err)
		return
	}
}
