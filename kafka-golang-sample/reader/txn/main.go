package main

import (
	"context"
	"kafka-golang-sample/common"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// err: unavailable when GroupID is not set
// 需要设置消费者组才可以 使用 commitMessages

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
		ctx := context.TODO()
		message, err := r.FetchMessage(ctx)
		if err != nil {
			log.Printf("read message err: %v\n", err)
			break
		}
		log.Printf("topic: %s, partitionId: %d, offset: %d, value: %v\n",
			message.Topic, message.Partition, message.Offset, string(message.Value))
		// TODO: logic here
		time.Sleep(1 * time.Second)
		// after logic then commit
		err = r.CommitMessages(ctx, message)
		if err != nil {
			log.Printf("commit message err: %v\n", err)
			return
		}
	}

	return

}
