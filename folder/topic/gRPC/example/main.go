package main

import (
	"os"
	"os/signal"
	"syscall"

	"example/app"
	"example/conf"
	"example/protodatasvr"
	"example/rpcclient"
	"example/rpcserver"
)

func main() {

	conf.NewConfiger("./app.conf")

	app.NewRPCSvr(conf.Configer.RpcConf.RpcPort)
	protodatasvr.RegisterEmailServiceServer(app.RPCSvr, rpcserver.SendEmailServer{})

	rpcclient.RpcClient()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quitChan
	app.StopRPCSvr()
}
