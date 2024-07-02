package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(h *Handler) *gin.Engine {
	r := gin.Default()

	// Ruta de ping
	r.GET("/ping", h.PingHandler)

	// Rutas de órdenes

	r.GET("/orders/:id", h.GetOrder)
	r.GET("/orders/user/:user_id", h.GetOrdersByUser)
	r.POST("/orders/", h.CreateOrder)
	r.PUT("/orders/:id/advance", h.AdvanceOrder)

	return r
}
