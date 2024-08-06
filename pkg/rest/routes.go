package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(h *Handler, userService UserService) *gin.Engine {
	r := gin.Default()

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// Middleware de autenticación
	authMiddleware := NewAuthMiddleware(userService)

	// Ruta de ping
	r.GET("/ping", h.PingHandler)

	// Rutas de órdenes protegidas
	auth := r.Group("/")
	auth.Use(authMiddleware.Middleware())
	{
		auth.GET("/order/:order_id", h.GetOrder)
		auth.GET("/order/user/:user_id", h.GetOrdersByUser)
		auth.GET("/order/professional/:professional_id", h.GetOrderByProfessional)
		auth.GET("/order/professional/:professional_id/day/:day", h.GetOrdersByProfessionalAndDay)
		auth.POST("/order", h.CreateOrder)
		auth.PUT("/order/:order_id/accept", h.AcceptOrder)
		auth.PUT("/order/:order_id/complete", h.CompleteOrder)
		auth.PUT("/order/:order_id/cancel", h.CancelOrder)
		auth.GET("/user/:user_id", h.GetUser)
		auth.GET("/professional/:professional_id", h.GetProfessional)
		auth.GET("/professionals/location/:department/:district", h.GetProfessionalsByLocation)
		auth.GET("/professionals/:department/:district/:profession", h.GetProfessionalsByLocationAndProfession)
	}

	// Rutas de registro y login
	r.GET("/locations", h.GetLocations)

	r.POST("/user", h.CreateUser)
	r.POST("/professional", h.CreateProfessional)
	r.POST("/user/login", h.LoginUser)
	r.POST("/professional/login", h.LoginProfessional)

	// Añadir en tu archivo de rutas
	r.POST("/logout", h.Logout)

	return r
}
