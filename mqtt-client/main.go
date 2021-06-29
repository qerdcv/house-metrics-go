package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func Run()  {
	log.Println("Running mqtt mqtt-client")
	opts := SetupOpts()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
