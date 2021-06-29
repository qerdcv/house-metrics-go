package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"house-metrics/mqtt-client"
	"log"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go client.Run()
	log.Fatal(http.ListenAndServe("localhost:3005", nil))
}
