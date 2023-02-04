package broker

import (
	"SiverPineValley/trailer-manager/config"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func InitBroker() (err error) {
	config := config.GetConfig()

	host := config.Broker.Host
	port := config.Broker.Port
	if host == "" {
		host = "0.0.0.0"
	}

	if port == "" {
		port = 1883
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("test")
	opts.SetPassword("test")
	opts.SetDefaultPublishHandler(broker.MessagePubHandler)
	opts.OnConnect = broker.ConnectHandler
	opts.OnConnectionLost = broker.ConnectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Disconnect(250)
}