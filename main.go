package main

import (
	"load_balancer/src/config"
	"load_balancer/src/routes"
	"load_balancer/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

func serverHandler(c *gin.Context) {
	type Input struct {
		Url string `json:"url"`
	}
	var url Input
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	newClient := config.Client {
		Url: url.Url,
	}

	config.JoinServer(newClient);
	c.JSON(200, gin.H{"message": "Joined successfully", "url": url.Url})
}

func MonitorServers () {
	for {
		services.AliveCheck();
		time.Sleep(30 * time.Second);
	}
}

func main() {
	go MonitorServers();
	r := gin.Default();
	r.POST("/join", serverHandler);
	serverGroup := r.Group("/api");
	routes.RegisterServerRoutes(serverGroup);
	r.Run(":8000");
}