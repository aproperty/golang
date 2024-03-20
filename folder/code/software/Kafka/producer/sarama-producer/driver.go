package main

import (
	"github.com/Shopify/sarama"

	"github.com/golang/glog"
)

var newProducer sarama.AsyncProducer

func InitAsynProducer(addrs []string) error {

	var err error

	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.ChannelBufferSize = 1 << 12 // config.ChannelBufferSize = 1 << 18
	config.Net.MaxOpenRequests = 32
	config.Producer.RequiredAcks = sarama.WaitForLocal // high level ack, Must ensure that data is not lost
	// config.Producer.Retry.Backoff = time.Duration(math.MaxInt32)
	config.Metadata.Full = true
	config.Producer.Retry.Max = 1 << 10

	newProducer, err = sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		glog.Errorln(err)
		// panic(err)
		return err
	}

	go chkAsyncProducer(newProducer)
	return nil
}

func chkAsyncProducer(p sarama.AsyncProducer) {
	errors := p.Errors()
	success := p.Successes()
	for {
		select {
		case err := <-errors:
			if err != nil {
				glog.Error("producer errror", err)
			}
		case <-success:
		}
	}
}
