package main

import (
	"context"
	"kafka-golang-sample/common"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{common.Address},
		Topic:     "job",
		Partition: 1,
		MinBytes:  1,
		MaxBytes:  1e6,
		MaxWait:   500 * time.Millisecond,
	})
	defer r.Close()

	for {
		message, err := r.ReadMessage(context.TODO())
		if err != nil {
			log.Printf("read message err: %v\n", err)
			break
		}
		log.Printf("topic: %s, partitionId: %d, offset: %d, value: %v\n",
			message.Topic, message.Partition, message.Offset, string(message.Value))
		time.Sleep(1 * time.Second)
	}

	return

}
