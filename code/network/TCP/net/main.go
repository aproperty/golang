package main

import (
	"fmt"
	"time"

	"temporary/tcpserver"
)

type DefaultIoHandlerFactory struct {
}

func (d *DefaultIoHandlerFactory) CreateIoHandler() tcpserver.IoHandler {
	return new(DefaultIoHandler)
}

func main() {

	server := tcpserver.GetServer()
	err := server.Start(new(DefaultIoHandlerFactory))
	if err != nil {
		fmt.Println("Server Start err: " + err.Error())
		return
	}

	for {
		time.Sleep(time.Second * 20)
	}
}
