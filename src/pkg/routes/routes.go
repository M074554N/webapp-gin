package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoutes(r *gin.Engine) {
	ping := r.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
