package logrustool

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {

	cfg := &LogConfig{
		FilePath:          "./log/",
		FileName:          "testlog",
		Level:             "Debug",
		Formatter:         "Text",
		LineLocateEnabled: true,
	}
	InitLog(cfg)

	log.WithFields(log.Fields{
		"fieldOne": "github.com/sirupsen/logrus",
		"fieldTwo": "github.com/lestrrat-go/file-rotatelogs",
	}).Info("Using logrus.")

	log.Info("start running")

	// quit := make(chan bool)
	// <-quit
}
