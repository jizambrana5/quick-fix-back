package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	domain "github.com/jizambrana5/quickfix-back/pkg/domain/auth"
	"github.com/jizambrana5/quickfix-back/pkg/utils"
)

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

type AuthService interface {
	Authenticate(username, password string) (string, error)
	ValidateToken(tokenString string) error
}

// PingHandler is a sample handler for the public route /api/ping
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

// LoginHandler handles user login and JWT token generation
func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Example authentication logic (replace with your actual authentication mechanism)
	if loginData.Username == "user1" && loginData.Password == "password1" {
		token, err := domain.GenerateJWT(loginData.Username, "user")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// AdminHandler handles admin-only access (protected endpoint)
func AdminHandler(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin access granted"})
}

// AuthMiddleware verifica y decodifica el token JWT en las rutas protegidas
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c.Request)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Verificar si el usuario existe en la base de datos
		user, err := repository.GetUserByUsername(claims.Username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
			c.Abort()
			return
		}

		// Verificar el rol del usuario
		if user.Role != claims.Role {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		// Establecer los valores en el contexto
		c.Set("user", user)
		c.Next()
	}
}

// extractToken extrae el token JWT del encabezado Authorization
func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}

// Routes configures application routes
func Routes(router *gin.Engine) {
	// Public routes
	public := router.Group("/api")
	{
		// Public endpoint for testing (no authentication required)
		public.GET("/ping", PingHandler)
		// Public endpoint for login (JWT token generation)
		public.POST("/login", LoginHandler)
	}

	// Middleware de autenticación JWT
	private := router.Group("/api/private")
	private.Use(AuthMiddleware()) // Middleware para verificar JWT
	{
		private.GET("/admin", AdminHandler)
	}
}
