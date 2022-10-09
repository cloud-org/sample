package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "job"
	partition := 1 // 分区

	ctx := context.TODO()
	address := "127.0.0.1:9092"
	// func DialLeader(
	// 	ctx context.Context,
	//	network string,
	//	address string,
	//	topic string,
	//	partition int)
	// (*Conn, error)
	conn, err := kafka.DialLeader(
		ctx, "tcp", address, topic, partition)
	if err != nil {
		log.Printf("conn err: %v\n", err)
		return
	}
	defer conn.Close()

	//fmt.Printf("conn is: %+v\n", conn)

	// write deadline: 10s 过后写会出现 io/timeout 然而看注释是也有可能写成功..
	// Even if write times out, it may return n > 0, indicating that some of the
	// data was successfully written.
	//_ = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	time.Sleep(3 * time.Second) // 调用 time.Sleep 睡眠 3s
	numbers, err := conn.WriteMessages(
		kafka.Message{Value: []byte("seven")},
		//kafka.Message{Value: []byte("two")},
		//kafka.Message{Value: []byte("three")},
	)
	if err != nil {
		log.Printf("write message err: %v\n", err)
		return
	}

	log.Printf("write numbers of bytes is: %v\n", numbers)

	return

}
