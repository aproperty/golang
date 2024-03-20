
`cd proto/`  

`protoc -I=/home/chan/Desktop/chanprogo/somemodule --go_out=             /home/chan/Desktop/chanprogo/somemodule              `
`protoc -I=/home/chan/Desktop/chanprogo/somemodule  `
`protoc                                            --go_out=plugins=grpc:../protodatasvr            *.proto`   
`protoc                                            --go-grpc_out=.                            proto/*.proto`   


