package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/jonatak/ble-security-mqtt/internal/haos"
	mqtthandler "github.com/jonatak/ble-security-mqtt/internal/mqttHandler"
)

func main() {
	err := godotenv.Load("/Users/billaudjonathan/Projects/perso/ble-security-mqtt/ble-security-mqtt/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mqttCon, err := mqtthandler.NewMqttContext()
	if err != nil {
		log.Fatal(err)
	}
	defer mqttCon.Close()

	if mqttCon != nil {
		log.Println("success")
	}

	if err := mqttCon.Subscribe(); err != nil {
		log.Fatal(err)
	}

	haosClient := haos.NewClient()

	state, err := haosClient.GetAlarmState()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Alarm state is %s\n", state)

	for c := range mqttCon.BleChan {
		fmt.Println(c)
	}
}
