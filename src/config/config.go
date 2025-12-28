package config

import "sync"

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

func GetCurrentRobin() string {
	SerArrMutex.Lock()
	defer SerArrMutex.Unlock()
	index = ((index + 1) % len(ServerArray));
	return ServerArray[index].Url;
}