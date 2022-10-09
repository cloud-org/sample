package main

import (
	"kafka-golang-sample/common"
	"log"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "example_create_topic"

	conn, err := kafka.Dial("tcp", common.Address)
	if err != nil {
		log.Printf("dial err: %v\n", err)
		return
	}

	defer conn.Close()

	// Controller requests kafka for the current controller and returns its URL
	// 获取 leader broker
	controller, err := conn.Controller()
	if err != nil {
		log.Printf("get current controller err: %v\n", err)
		return
	}
	log.Printf("current controller is %+v\n", controller)
	controllerConn, err := kafka.Dial(
		"tcp",
		net.JoinHostPort(
			controller.Host,
			strconv.Itoa(controller.Port),
		),
	)

	if err != nil {
		log.Printf("dial controller err: %v\n", err)
		return
	}
	defer controllerConn.Close()

	//CreateTopics 可以传多个 不过正常给业务用不会这么玩的
	err = controllerConn.CreateTopics(
		kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     10,
			ReplicationFactor: 1,
		},
		kafka.TopicConfig{
			Topic: "unset_partition_topic",
			// topic is unset_partition_topic, partitions count is 1
			NumPartitions:     -1, // -1 indicates unset 不过貌似最终也是 1
			ReplicationFactor: 1,
		},
	)
	if err != nil {
		log.Printf("create topic err: %v\n", err)
		return
	}

	log.Printf("create success")
}
