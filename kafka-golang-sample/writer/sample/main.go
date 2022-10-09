package main

import (
	"context"
	"kafka-golang-sample/common"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

//PartitionBalancer 自定义分区选择器，满足业务需要
type PartitionBalancer struct {
}

func (p *PartitionBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	log.Printf("msg partition: %v, value: %v\n", msg.Partition, string(msg.Value))
	return msg.Partition
}

func main() {
	writer := kafka.Writer{
		Addr:         kafka.TCP(common.Address),
		Topic:        "job",
		Balancer:     &PartitionBalancer{},
		BatchSize:    100,
		BatchTimeout: 500 * time.Millisecond, // 500ms 提交一次
	}

	err := writer.WriteMessages(
		context.TODO(),
		kafka.Message{
			Partition: 1, // 正常 write message 不应该指定
			Key:       nil,
			Value:     []byte("ashing"),
		},
		kafka.Message{
			Partition: 2,
			Key:       nil,
			Value:     []byte("xzx"),
		},
	)
	if err != nil {
		log.Printf("write message: %v\n", err)
		return
	}

	return

}
