package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	jsoniter "github.com/json-iterator/go"

	"github.com/golang/glog"
)

func SendOrderToMw(v interface{}, groupID string) error {

	buf, err := jsoniter.Marshal(v)
	if err != nil {
		glog.Error(err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: fmt.Sprintf("market_group_%s", groupID),
		Value: sarama.ByteEncoder(buf),
	}

	newProducer.Input() <- msg
	return nil
}

func SendMakeVolToMw(v interface{}, groupID string) error {

	buf, err := jsoniter.Marshal(v)
	if err != nil {
		glog.Error(err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Key:   sarama.ByteEncoder([]byte("make_volume")),
		Topic: fmt.Sprintf("make_vol_%s", groupID),
		Value: sarama.ByteEncoder(buf),
	}

	newProducer.Input() <- msg
	return nil
}
