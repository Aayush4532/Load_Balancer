package routes

import (
	"load_balancer/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterServerRoutes(r *gin.RouterGroup) {
	r.GET("/api", controllers.ServerApiHandler)
}
