package config

import (
	"fmt"
	"sync"
)

type Client struct {
	Url string
}

var (
	ServerArray []Client;
	index int = 0;
	SerArrMutex sync.Mutex;
)

func JoinServer(client Client) {
	SerArrMutex.Lock()
	defer SerArrMutex.Unlock()
	ServerArray = append(ServerArray, client);
}

func GetCurrentRobin() (string, error) {
	SerArrMutex.Lock()
	defer SerArrMutex.Unlock()
	if len(ServerArray) == 0 {
		return "", fmt.Errorf("no servers available")
	}
	index = ((index + 1) % len(ServerArray));
	return ServerArray[index].Url, nil;
}