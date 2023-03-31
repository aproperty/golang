package mqttx

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewTlsConfig() *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../config/cert/rootCA.pem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	return &tls.Config{
		RootCAs: certpool,
	}
}

var client gomqtt.Client // 跟 MQTT 服务器连接的实例

func Start(broker, clientID string) {

	// gomqtt.DEBUG = log.New(os.Stdout, "", 0)
	// gomqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := gomqtt.NewClientOptions()

	// opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)

	opts.AddBroker(broker)

	if len(clientID) > 0 {
		opts.SetClientID(clientID)
	}

	// opts.SetUsername("emqx")
	// opts.SetPassword("mypassword")

	opts.SetOnConnectHandler(onConnectHandler(opts.OnConnect))
	opts.SetConnectionLostHandler(onConnectionLostHandler(opts.OnConnectionLost))

	// 设置消息回调处理函数
	var f gomqtt.MessageHandler = func(client gomqtt.Client, msg gomqtt.Message) {
		fmt.Printf("fmt - SetDefaultPublishHandler: TOPIC: %s\n", msg.Topic())
		fmt.Printf("fmt - SetDefaultPublishHandler: MSG: %s\n", msg.Payload())
	}
	opts.SetDefaultPublishHandler(f)
	// opts.SetPingTimeout(1 * time.Second)

	tlsConfig := NewTlsConfig()
	opts.SetTLSConfig(tlsConfig)

	client = gomqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// 连接上服务器的操作
func onConnectHandler(handler gomqtt.OnConnectHandler) gomqtt.OnConnectHandler {
	return func(c gomqtt.Client) {
		for _, item := range subscribers {
			subscribe(item)
		}
		// handler(c)
	}
}

// 丢失连接的操作(自动重连)
func onConnectionLostHandler(handler gomqtt.ConnectionLostHandler) gomqtt.ConnectionLostHandler {
	return func(c gomqtt.Client, e error) {
		fmt.Printf("Connect lost: %v", e)
		// handler(c, e)
	}
}
