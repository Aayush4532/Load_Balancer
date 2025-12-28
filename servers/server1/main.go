package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Server 1",
	})
}

func registerWithLoadBalancer() {
	data := map[string]string{
		"url": "http://localhost:5171",
	}

	body, _ := json.Marshal(data)

	resp, err := http.Post(
		"http://localhost:8000/join",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("Failed to register with LB:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Server 1 registered with load balancer")
}

func pinghandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong from Server 1",
	})
}

func main() {
	r := gin.Default()
	r.GET("/api", handler)
	r.GET("/ping", pinghandler);

	log.Println("Server 1 running on :5171")
	registerWithLoadBalancer()
	r.Run(":5171")
}
