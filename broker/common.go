package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(client mqtt.Client, topic string) {
	text := fmt.Sprintf("Message %d")
	token := client.Publish(topic, 0, false, text)
	token.Wait()
}

func Subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
