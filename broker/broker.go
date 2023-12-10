package broker

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	tm "github.com/trailer-manager/trailer-manager-common"
	"github.com/trailer-manager/trailer-manager-common/config"
	tmError "github.com/trailer-manager/trailer-manager-common/error"
	"github.com/trailer-manager/trailer-manager-common/logger"
	"github.com/trailer-manager/trailer-manager-fluent/model/api"
	"time"
)

type Broker struct {
	qos     byte
	handler mqtt.MessageHandler
}

var (
	topicFuncMap = map[string]Broker{
		"gps/log": {1, GpsLogHandler},
	}
)

func initTopic(client mqtt.Client) (err error) {
	config := config.GetConfig()
	topics := config.Broker.Topics

	for _, t := range topics {
		token := client.Subscribe(t, topicFuncMap[t].qos, topicFuncMap[t].handler)
		token.Wait()
		fmt.Printf("Subscribed to topic %s", t)
	}

	return
}

func InitBroker() (err error) {
	config := config.GetConfig()

	host := config.Broker.Host
	port := config.Broker.Port
	clientId := config.Broker.ClientId
	username := config.Broker.Username
	pwd := config.Broker.Password

	if host == "" {
		host = tm.HostDefault
	}

	if port <= 0 {
		port = tm.PortDefault
	}

	if clientId == "" {
		return tmError.ConfigErr
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(pwd)
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.OnConnect = ConnectHandler
	opts.OnConnectionLost = ConnectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Fatal(token.Error().Error())
	}
	initTopic(client)
	//client.Disconnect(250)
	return nil
}

func PublishTest(client mqtt.Client) {
	t := model.GpsLogRequest{
		Sid:     "e6:61:64:8:43:78:68:23",
		Lat:     "",
		Lon:     "",
		Speed:   "0",
		WifiLoc: []string{"csg"},
		Battery: 100,
	}
	text, _ := json.Marshal(t)
	token := client.Publish("gps/log", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}
