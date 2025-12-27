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
		"message": "Hello from Server 4",
	})
}

func registerWithLoadBalancer() {
	data := map[string]string{
		"url": "http://localhost:5174",
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

	log.Println("Server 4 registered with load balancer")
}

func main() {
	r := gin.Default()
	r.GET("/api", handler)

	log.Println("Server 4 running on :5174")
	registerWithLoadBalancer()
	r.Run(":5174")
}
