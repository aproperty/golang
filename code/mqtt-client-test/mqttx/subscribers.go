package mqttx

import (
	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeType struct {
	Topic    string
	Qos      byte
	Callback gomqtt.MessageHandler

	RetryTimes int // 表示订阅失败后的重试次数，如果为0，则表示一直重试下去。
}

var subscribers = make([]SubscribeType, 0)

// 注册订阅消息
func AppendToSubscribers(item SubscribeType) {
	subscribers = append(subscribers, item)
}
