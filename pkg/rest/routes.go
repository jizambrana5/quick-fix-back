package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(h *Handler) *gin.Engine {
	r := gin.Default()

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// Ruta de ping
	r.GET("/ping", h.PingHandler)

	// Rutas de órdenes
	r.GET("/order/:order_id", h.GetOrder)
	r.GET("/order/user/:user_id", h.GetOrdersByUser)
	r.GET("/order/professional/:professional_id", h.GetOrderByProfessional)
	r.GET("/order/professional/:professional_id/day/:day", h.GetOrdersByProfessionalAndDay)

	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:order_id/accept", h.AcceptOrder)
	r.PUT("/order/:order_id/complete", h.CompleteOrder)
	r.PUT("/order/:order_id/cancel", h.CancelOrder)

	// Rutas de registro
	r.POST("/user", h.CreateUser)
	r.GET("/user/:user_id", h.GetUser)
	r.POST("/professional", h.CreateProfessional)
	r.GET("/professional/:professional_id", h.GetProfessional)
	r.GET("/professionals/location/:department/:district", h.GetProfessionalsByLocation)
	r.GET("/professionals/:department/:district/:profession", h.GetProfessionalsByLocationAndProfession)

	// Ruta para obtener ubicaciones
	r.GET("/locations", h.GetLocations)

	// Login
	r.POST("/user/login", h.LoginUser)
	r.POST("/professional/login", h.LoginProfessional)

	return r
}
