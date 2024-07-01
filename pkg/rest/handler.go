package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler maneja la solicitud GET /ping
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
