package main

import (
	"fmt"
	"temporary/mqttx"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	testMessageHandler := func(client gomqtt.Client, msg gomqtt.Message) {
		fmt.Printf("fmt - Subscribe - TOPIC: %s\n", msg.Topic())
		fmt.Printf("fmt - Subscribe - MSG:   %s\n", msg.Payload())
	}
	item := mqttx.SubscribeType{
		Topic:      "topic_test",
		Qos:        byte(2),
		Callback:   testMessageHandler,
		RetryTimes: 0,
	}
	mqttx.AppendToSubscribers(item) // 注册需要订阅的消息，连接后会自动订阅

	var broker = "localhost"
	// var port = 1883
	// server := fmt.Sprintf("tcp://%s:%d", broker, port)
	var port = 8883
	server := fmt.Sprintf("tls://%s:%d", broker, port)

	clientID := "go_mqtt_client"
	mqttx.Start(server, clientID)

	quit := make(chan bool)
	<-quit
}
