package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"strings"

	"github.com/Shopify/sarama"
)

// Info ...
type Info struct {
	ID     int    `json:"myID"`
	Name   string `json:"myName"`
	Time   string `json:"myTime"`
	Enable bool   `json:"myEnable"`
}

// PushStrings ...
var PushStrings chan string

func main() {
	PushStrings = make(chan string, 1000)

	for i := 400; i < 500; i++ {
		var tmp Info
		tmp.ID = i
		tmp.Name = "name-" + strconv.Itoa(i)
		tmp.Time = strconv.FormatInt(time.Now().Unix(), 10)
		tmp.Enable = true

		tjson, err := json.Marshal(tmp)
		if err != nil {
			panic(err)
		}

		PushStrings <- string(tjson)
		fmt.Printf("%v - ", i)
	}

	go PushStringDataToKafka(PushStrings,
		"topic001",
		"120.79.56.54:9092,112.74.173.118:9092,47.106.188.16:9092")

	quit := make(chan bool)
	<-quit
}

// PushStringDataToKafka ...
func PushStringDataToKafka(c chan string, topic string, connect string) {

	producer, err := sarama.NewAsyncProducer(strings.Split(connect, ","), nil)
	if err != nil {
		fmt.Println("producer NewAsyncProducer err: ", err.Error())
		panic(err)
	}

	fmt.Println("PushStringDataToKafka begin...")
	count := 0
	for {
		fmt.Print(" -  ")
		dataSave := <-c
		count++
		fmt.Printf("count:%v \n", count)

		producerMessage := &sarama.ProducerMessage{}
		producerMessage.Topic = topic
		producerMessage.Value = sarama.ByteEncoder([]byte(dataSave))

		select {
		case producer.Input() <- producerMessage:

		case produceSendErr := <-producer.Errors():
			fmt.Println("Failed to produce message", produceSendErr)
		}
	}
}
