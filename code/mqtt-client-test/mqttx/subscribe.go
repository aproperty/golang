package mqttx

import (
	"time"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

func subscribe(item SubscribeType) {
	times := 0
	for {
		token, err := subscribeItem(item)
		if err != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			}
			panic(err)
		}
		if token.Wait() && token.Error() != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			}
			panic(token.Error())
		}
		break
	}
}

func subscribeItem(item SubscribeType) (token gomqtt.Token, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	// token := client.Subscribe("testtopic/#", 0, nil)
	token = client.Subscribe(item.Topic, item.Qos, item.Callback)
	return
}

// 取消订阅
// if token := client.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
// 	os.Exit(1)
// }
