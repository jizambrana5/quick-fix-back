package domain

import (
	"github.com/golang-jwt/jwt"

	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

var jwtKey = []byte("your_jwt_secret_key")

type Service struct{}

// Claims estructura para el JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// NewService devuelve una implementación de UserService
func NewService() *Service {
	return &Service{}
}

var _ rest.AuthService = (*Service)(nil)
