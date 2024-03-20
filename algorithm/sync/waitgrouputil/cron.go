package waitgrouputil

import (
	"time"
)

var IsQuit bool

type CronObj struct {
	lastTime int64
	// lock     sync.RWMutex
}

func (cronObj *CronObj) Init(intervalSec int64, f func()) {

	intervalSec *= int64(time.Second)

	handler := func(...interface{}) {
		for {
			if IsQuit {
				break
			}

			nowTime := time.Now().UnixNano()

			if cronObj.lastTime+intervalSec < nowTime {
				cronObj.lastTime = nowTime
				f()
			}

			time.Sleep(time.Millisecond * 100)
		}
	}

	WaitGroup.WrapGoroutine(handler)
}
