package config

import "sync"

var (
	num      = 0
	numMutex sync.Mutex
	Robin = num;
	RobinMutex sync.Mutex
)

func GetId() int {
	numMutex.Lock()
	defer numMutex.Unlock()
	num += 1
	return num
}

type Client struct {
	Url string
}

var (
	Room      = make(map[int]Client)
	RoomMutex sync.Mutex
)

func JoinServer(client Client) {
	RoomMutex.Lock()
	defer RoomMutex.Unlock()
	id := GetId()
	Room[id] = client
}

func GetCurrentRobin() int {
	RobinMutex.Lock()
	defer RobinMutex.Unlock()
	Robin = ((Robin + 1) % num);
	if(Robin == 0) {
		Robin = num;
	}
	return Robin
}