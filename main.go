package main

import (
	client "house-metrics/mqtt-client"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go client.Run()
	log.Fatal(http.ListenAndServe("localhost:3005", nil))
}
