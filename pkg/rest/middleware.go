package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	userService UserService
}

func NewAuthMiddleware(userService UserService) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
	}
}

func (a *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		session, err := a.userService.ValidateSession(c.Request.Context(), token)
		if err != nil || session.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", session.UserID)
		c.Next()
	}
}
