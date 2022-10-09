package main

import (
	"fmt"
	"kafka-golang-sample/common"
	"log"

	"github.com/segmentio/kafka-go"
)

// for list topics

func main() {
	conn, err := kafka.Dial("tcp", common.Address)
	if err != nil {
		log.Printf("dial err: %v\n", err)
		return
	}

	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Printf("read partitions err: %v\n", err)
		return
	}

	topicMap := make(map[string][]int)
	for i := 0; i < len(partitions); i++ {
		//fmt.Printf("partition%d is %+v\n", i, partitions[i])
		//fmt.Printf(
		//	"topic: %s, partitionId: %d\n",
		//	partitions[i].Topic,
		//	partitions[i].ID,
		//)
		topicMap[partitions[i].Topic] = append(topicMap[partitions[i].Topic], partitions[i].ID)
	}

	for k, v := range topicMap {
		fmt.Printf("topic is %v, partitions count is %v\n", k, len(v))
	}
	//output:
	//topic is __consumer_offsets, partitions count is 50
	//topic is job, partitions count is 2000
	//topic is sun, partitions count is 1

	return
}
