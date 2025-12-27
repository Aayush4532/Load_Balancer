package controllers

import (
	"io"
	"load_balancer/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerApiHandler(c *gin.Context) {
	num := config.GetCurrentRobin()
	url := config.Room[num].Url + "/api"

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to reach server"})
		return
	}
	defer resp.Body.Close()
	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
