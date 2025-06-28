package mqtthandler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttContext struct {
	Client mqtt.Client
}

func (m *MqttContext) Close() {
	m.Client.Disconnect(100)
}

func NewMqttContext() (*MqttContext, error) {

	var port int

	broker := os.Getenv("MQTT_HOST")
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	clientId := os.Getenv("MQTT_CLIENT_ID")
	port, err := strconv.Atoi(os.Getenv("MQTT_PORT"))

	if err != nil {
		return nil, errors.New("invalid MQTT_PORT")
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	}

	return &MqttContext{
		Client: client,
	}, nil
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}
