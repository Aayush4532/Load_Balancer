package routes

import (
	"load_balancer/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterServerRoutes(r *gin.RouterGroup) {
	r.Any("/*proxy", controllers.ServerApiHandler)
}
