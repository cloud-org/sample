package common

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const (
	Topic   = "job"
	Address = "127.0.0.1:9092"
)

//GetLeaderConn 因为是 demo 所以写死, 之后可以新增其他的入参
//partition int 分区 ID
func GetLeaderConn(partition int) (*kafka.Conn, error) {

	topic := Topic
	ctx := context.TODO()
	address := Address

	// func DialLeader(
	// 	ctx context.Context,
	//	network string,
	//	address string,
	//	topic string,
	//	partition int)
	// (*Conn, error)
	return kafka.DialLeader(ctx, "tcp", address, topic, partition)
}
