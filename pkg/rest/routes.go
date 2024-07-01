package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Rutas públicas
	public := router.Group("/api")

	// Rutas privadas (requieren autenticación, etc.)

	// Ruta de ping
	public.GET("/ping", PingHandler)
}
