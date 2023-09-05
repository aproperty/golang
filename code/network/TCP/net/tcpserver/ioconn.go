package tcpserver

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// type IoConn interface {
// 	Start(IoHandler, net.Conn, uint32)
// 	Close()

// 	Read(uint32)
// 	Write([]byte)

// 	RemoteName() string
// 	RemoteAddr() string

// 	OnClose(**string)
// }

type MyTCPConn struct {
	conn    net.Conn
	handler IoHandler

	readedData    []byte
	readedDataLen int
}

func (c *MyTCPConn) Start(handler IoHandler, conn net.Conn, maxSendQueue uint32) {

	go func() {

		c.conn = conn
		c.handler = handler

		c.readedData = make([]byte, 7000)
		c.readedDataLen = 0

		var myKey *string

		var stopLoop bool
		for !stopLoop {
			// fmt.Println("Read data begin. address:" + conn.RemoteAddr().String() + "\n")

			if err := c.conn.SetReadDeadline(time.Now().Add(time.Minute * (time.Duration(5)))); err != nil {
				fmt.Println("set read timeout fail, addr:" + conn.RemoteAddr().String())
				stopLoop = true
			} else {

				readBytes, err := c.conn.Read(c.readedData[c.readedDataLen:])
				if err != nil {
					fmt.Println("Read data fail. err:" + err.Error() + ", addr:" + conn.RemoteAddr().String())
					stopLoop = true
				} else if readBytes == 0 {
					fmt.Println("Read nothing.., addr:" + conn.RemoteAddr().String())
					stopLoop = true
				} else {
					fmt.Println()
					fmt.Println("Receive data len:" + strconv.Itoa(readBytes))
					c.readedDataLen += readBytes

					for c.readedDataLen > 0 {

						suc, msgLen := c.handler.OnReadFinished(&myKey, conn, c.readedData[:c.readedDataLen])
						if !suc {
							fmt.Println("OnReadFinished fail.")
							stopLoop = true
							break

						} else if msgLen > 0 {

							{
								leftLen := c.readedDataLen - msgLen
								for i := 0; i < leftLen; i++ {
									c.readedData[i] = c.readedData[i+msgLen]
								}
							}

							c.readedDataLen -= msgLen

						} else {
							fmt.Println("No complete msg left.")
							break
						}
					}
				}
			}
		} //'for !stopLoop'

		c.OnClose(&myKey)

	}() //End of 'go func'
}

// Close ...
func (oTConn *MyTCPConn) Close() {
}

func (oTConn *MyTCPConn) Read(bytes uint32) {
}

func (oTConn *MyTCPConn) Write(data []byte) {
}

// RemoteName ...
func (oTConn *MyTCPConn) RemoteName() string {
	return oTConn.conn.RemoteAddr().Network()
}

// RemoteAddr ...
func (oTConn *MyTCPConn) RemoteAddr() string {
	return oTConn.conn.RemoteAddr().String()
}

// OnClose ...
func (c *MyTCPConn) OnClose(myKey **string) {
	fmt.Println("ObdTcpConn OnClose begin. addr:" + c.conn.RemoteAddr().String())

	GetServer().OnIoDisCon(*myKey)
	c.conn.Close()
	c.handler.OnClosed()

	fmt.Println("MyTcpConn OnClose finish..")
}
