package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

// Info ...
type Info struct {
	ID     int    `json:"myID"`
	Name   string `json:"myName"`
	Time   string `json:"myTime"`
	Enable bool   `json:"myEnable"`
}

// ProducerAnalisysChans 解析消息chan
var ProducerAnalisysChans []chan *Info

func main() {
	ProducerAnalisysChans = make([]chan *Info, 10)   // 一个数组，其中每个元素都是通道
	anaProducers := make([]sarama.AsyncProducer, 10) // 一个数组

	var i uint32
	for i = 0; i < 10; i++ {
		ProducerAnalisysChans[i] = make(chan *Info, 10000)

		var err error
		anaProducers[i], err = sarama.NewAsyncProducer(
			strings.Split("39.104.81.59:9092,39.104.153.163:9092,39.104.145.49:9092", ","), nil)
		if err != nil {
			panic(err)
		}
	}

	for i = 0; i < 10; i++ {
		go InputAnalisysKafka("analysis-msg-topic", anaProducers[i], i)
	}

	defer func() {
		for i = 0; i < 10; i++ {
			if err := anaProducers[i].Close(); err != nil {
				fmt.Println("producer_ana[", i, "] Error closing the producer", err)
			}
			close(ProducerAnalisysChans[i])
		}
	}()

	for {
		time.Sleep(time.Second * 20)
	}
}
