package services

import (
	"load_balancer/src/config"
	"net/http"
	"time"
)

func PingHandler(client config.Client, index int) {
	beforeTime := time.Now()
	resp, err := http.Get(client.Url + "/ping")
	afterTime := time.Now()
	Latency := afterTime.Sub(beforeTime)
	if err != nil || resp.StatusCode != 200 {
		config.SerArrMutex.Lock()
		config.ServerArray = append(config.ServerArray[:index], config.ServerArray[index+1:]...)
		config.SerArrMutex.Unlock()
	} else {
		if Latency >= (2 * time.Second) {
			config.SerArrMutex.Lock()
			config.ServerArray = append(config.ServerArray[:index], config.ServerArray[index+1:]...)
			config.SerArrMutex.Unlock()
		}
	}
}

func AliveCheck() {
	ClientArray := config.ServerArray
	config.SerArrMutex.Lock()
	defer config.SerArrMutex.Unlock()
	for client := range ClientArray {
		go PingHandler(ClientArray[client], client)
	}
}
