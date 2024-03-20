package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Shopify/sarama"
)

// InputAnalisysKafka ...
func InputAnalisysKafka(topicName string, producer sarama.AsyncProducer, i uint32) {

	var count int
	for info := range ProducerAnalisysChans[i] {

		count++

		producerMessage := &sarama.ProducerMessage{}

		producerMessage.Topic = topicName
		producerMessage.Key = sarama.StringEncoder("ABCD999" + strconv.Itoa(count))

		encodeMsg, err := json.Marshal(info)
		if err != nil {
			fmt.Println("Producer_ana[" + strconv.Itoa(int(i)) + "] Encode err: " + err.Error())
		}
		producerMessage.Value = sarama.ByteEncoder(encodeMsg)

		select {
		case producer.Input() <- producerMessage:
			fmt.Println("Producer_ana[" + strconv.Itoa(int(i)) + "] message send success")

		case err := <-producer.Errors():
			fmt.Println("Failed to producer_ana[" + strconv.Itoa(int(i)) + "] message " + err.Error())
		}
	}

}
