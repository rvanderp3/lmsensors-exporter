package main

import (
	"fmt"
	_ "github.com/rvanderp3/prom-lm-sensors/pkg"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting Prometheus Exporter for lmsensors on port 9440")
	err := http.ListenAndServe(":9440", nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to start Prometheus metrics server: %v", err))
	}
}
