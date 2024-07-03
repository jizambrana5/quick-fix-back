package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(h *Handler) *gin.Engine {
	r := gin.Default()

	// Ruta de ping
	r.GET("/ping", h.PingHandler)

	// Rutas de órdenes
	r.GET("/orders/:order_id", h.GetOrder)
	r.GET("/orders/user/:user_id", h.GetOrdersByUser)
	r.GET("/orders/professional/:professional_id", h.GetOrderByProfessional)
	r.POST("/orders/", h.CreateOrder)
	r.PUT("/orders/:order_id/accept", h.AcceptOrder)
	r.PUT("/orders/:order_id/complete", h.CompleteOrder)
	r.PUT("/orders/:order_id/cancel", h.CancelOrder)

	// Rutas de registro
	r.POST("/user/", h.CreateUser)
	r.GET("/user/:user_id", h.GetUser)
	r.POST("/professional/", h.CreateProfessional)
	r.GET("/professional/:professional_id", h.GetProfessional)
	r.GET("/professionals/:department/:district", h.GetProfessionalsByLocation)

	// Ruta para obtener ubicaciones
	r.GET("/locations", h.GetLocations)

	return r
}
