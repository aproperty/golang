package clusterconsumer

import (
	"fmt"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
)

var consumer *cluster.Consumer

func KafkaInit(addrs []string, groupID string, topicId string, superReboot bool) error {

	topic := fmt.Sprintf("order_resp_%s", topicId)

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	if superReboot {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	var err error
	consumer, err = cluster.NewConsumer(addrs, groupID, []string{topic}, config)
	if err != nil {
		glog.Errorln(err)
		return err
	}

	go checkkConsumer(consumer)
	return nil
}

func checkkConsumer(c *cluster.Consumer) {
	errors := c.Errors()
	notiChan := c.Notifications()
	for {
		select {
		case err := <-errors:
			if err != nil {
				glog.Error("consumer error", err)
			}
		case <-notiChan:
		}
	}
}

func Close() {
	consumer.Close()
}

func GetMessagesChan() <-chan *sarama.ConsumerMessage {
	return consumer.Messages()
}
