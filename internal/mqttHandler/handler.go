package mqtthandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttContext struct {
	client   mqtt.Client
	espTopic string
	BleChan  chan string
}

func (m *MqttContext) Close() {
	m.client.Disconnect(100)
}

func (m *MqttContext) messageHandler(_ mqtt.Client, msg mqtt.Message) {
	var j Message
	if err := json.Unmarshal(msg.Payload(), &j); err != nil {
		log.Println("could not unmarshal msg")
	}
	m.BleChan <- fmt.Sprintf("Received message: %s from topic: %s\n", j, msg.Topic())
}

func NewMqttContext() (*MqttContext, error) {

	var port int
	var m *MqttContext = &MqttContext{}

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
	opts.SetDefaultPublishHandler(m.messageHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	}
	m.client = client
	m.espTopic = os.Getenv("MQTT_ESP_TOPIC")
	m.BleChan = make(chan string, 10)
	return m, nil
}

func (m *MqttContext) Subscribe() error {
	if token := m.client.Subscribe(m.espTopic, byte(0), nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}
