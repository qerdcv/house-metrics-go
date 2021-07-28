package client

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Run() {
	log.Println("Running mqtt mqtt-client")
	opts := SetupOpts()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

}
