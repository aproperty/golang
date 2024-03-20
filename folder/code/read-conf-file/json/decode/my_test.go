package jdconf

import (
	"runtime"
	"testing"
)

func TestReadConfig(t *testing.T) {

	if err := ReadConfig(); err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%+v \n", Conf)

	if Conf.Threads > 0 {
		runtime.GOMAXPROCS(Conf.Threads)
		t.Logf("Running with %v threads", Conf.Threads)
	}

	// quit := make(chan bool)
	// <-quit
}
