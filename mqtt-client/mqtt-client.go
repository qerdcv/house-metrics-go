package client

import (
	"log"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	temperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "temperature",
		Help: "Temperature in room",
	})
	humidity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "humidity",
		Help: "Humidity in room",
	})
	methane = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "methane",
		Help: "Methane level in the air",
	})
	propane = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "propane",
		Help: "Propane level in the air",
	})
	smoke = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "smoke",
		Help: "Smoke level in the air",
	})
)

var handlers = map[string]func(float64){
	"smoke":       smoke.Set,
	"methane":     methane.Set,
	"propane":     propane.Set,
	"temperature": temperature.Set,
	"humidity":    humidity.Set,
}

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	val, err := strconv.ParseFloat(string(msg.Payload()), 64)
	if err != nil {
		log.Printf("Failed to convert from topic %s", msg.Topic())
		return
	}
	handlers[msg.Topic()](val)
}

var connectionHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to mqtt")
	if token := client.SubscribeMultiple(map[string]byte{
		"temperature": 0,
		"humidity":    0,
		"propane":     0,
		"methane":     0,
		"smoke":       0,
	}, messageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func SetupOpts() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(":1883")
	opts.SetClientID("go_client")
	opts.OnConnect = connectionHandler
	return opts
}
