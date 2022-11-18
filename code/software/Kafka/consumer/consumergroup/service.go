package main

import (
	"fmt"
	"time"
)

// MainChannel ...
var MainChannel chan Info

func goAnalysis() {

	timeout := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(2))
			timeout <- true
		}
	}()

	var num int
	for {
		var msg Info
		select {

		case msg = <-MainChannel:
			num++
			fmt.Printf("%v goAnalysis --------- %+v \n", num, msg)

		case <-timeout:
			fmt.Println("-------")

		}
	}
}
