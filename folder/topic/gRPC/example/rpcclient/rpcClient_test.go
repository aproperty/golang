package rpcclient

import (
	"testing"

	"example/app"
	"example/conf"
	"example/protodatasvr"
	"example/rpcserver"
)

func init() {
	conf.NewConfiger("../app.conf")

	app.NewRPCSvr(conf.Configer.RpcConf.RpcPort)
	protodatasvr.RegisterEmailServiceServer(app.RPCSvr, rpcserver.SendEmailServer{})
}

func TestRpcClient(t *testing.T) {
	RpcClient()
}
