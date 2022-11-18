package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"
)

// RecvMessage ...
func RecvMessage(messageValue []byte) {

	anlmsg := &Info{}

	if err := json.Unmarshal(messageValue, anlmsg); err != nil {
		fmt.Println("Unmarshal faild. ", err.Error())
		return
	}

	CacheChannel <- *anlmsg
	fmt.Println(" - ")
}

func consume() {

	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	var zookeeperNodes []string
	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString("120.79.56.54:2181,112.74.173.118:2181,47.106.188.16:2181")

	// topics := strings.Split("topic1,topic2", ",")
	topics := []string{"topic001"}

	consumer, err := consumergroup.JoinConsumerGroup(
		"my-consumer-group-001",
		topics,
		zookeeperNodes,
		config)
	if err != nil {
		log.Fatalln(err)
	}

	for message := range consumer.Messages() {
		RecvMessage(message.Value)
		consumer.CommitUpto(message)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
		}
		close(MainChannel)
	}()
}
