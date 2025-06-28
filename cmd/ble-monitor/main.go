package main

import (
	"log"

	"github.com/joho/godotenv"
	mqtthandler "github.com/jonatak/ble-security-mqtt/internal/mqttHandler"
)

func main() {
	err := godotenv.Load()
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
}
