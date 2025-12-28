package controllers

import (
	"io"
	"load_balancer/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerApiHandler(c *gin.Context) {
	method := c.Request.Method
	proxyPath := c.Param("proxy")
	query := c.Request.URL.RawQuery
	body := c.Request.Body

	base, err := config.GetCurrentRobin()
	if err != nil {
		c.JSON(500, gin.H{"error": "it's not you, it's us, comeback later"})
		return
	}

	targetURL := base + "/api" + proxyPath
	if query != "" {
		targetURL += "?" + query
	}

	req, err := http.NewRequest(method, targetURL, body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create request"})
		return
	}

	req.Header = c.Request.Header.Clone()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "backend down"})
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		for _, val := range v {
			c.Writer.Header().Add(k, val)
		}
	}

	io.Copy(c.Writer, resp.Body)
}
