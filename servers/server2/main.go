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
		"message": "Hello from Server 2",
	})
}

func registerWithLoadBalancer() {
	data := map[string]string{
		"url": "http://localhost:5172",
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

	log.Println("Server 2 registered with load balancer")
}

func main() {
	r := gin.Default()
	r.GET("/", handler)

	log.Println("Server 2 running on :5172")
	registerWithLoadBalancer()
	r.Run(":5172")
}
