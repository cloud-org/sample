package main

import (
	"context"
	"fmt"
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

	//_ = conn.SetReadDeadline(time.Time{}) // 默认应该不设置就是零值
	//time.Sleep(3 * time.Second)
	first, err := conn.ReadFirstOffset()
	if err != nil {
		log.Printf("read first offset err: %v\n", err)
		return
	}
	log.Printf("first offset is %v\n", first)
	last, err := conn.ReadLastOffset()
	if err != nil {
		log.Printf("read last offset err: %v\n", err)
		return
	}
	log.Printf("last offset is %v\n", last)

	newOffset, err := conn.Seek(last, kafka.SeekAbsolute)
	if err != nil {
		log.Printf("seek err: %v\n", err)
		return
	}
	log.Printf("newOffset is %v\n", newOffset)

	for {
		message, err := conn.ReadMessage(1e6) // 如果读取不到消息，会阻塞
		if err != nil {
			log.Printf("read message err: %v\n", err)
			//continue // TODO: 可以考虑重新连接
			return
		}
		fmt.Printf("message Value: %s, Offset: %d, Time: %v\n", message.Value, message.Offset, message.Time)
		time.Sleep(1 * time.Second)
	}
	//batch := conn.ReadBatch(10e3, 1e6)
	//defer batch.Close()
	//for {
	//	message, err := batch.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	fmt.Printf("message is %+v\n", message)
	//	fmt.Printf()
	//}
}
