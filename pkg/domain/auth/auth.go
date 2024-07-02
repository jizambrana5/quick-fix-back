package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate verifica las credenciales del usuario y genera un token JWT si son válidas
func (s Service) Authenticate(username, password string) (string, error) {
	// Ejemplo de autenticación básica con usuarios en memoria o en base de datos
	// Aquí deberías implementar la lógica para verificar las credenciales en tu sistema de almacenamiento

	// Supongamos que tenemos un usuario en una estructura de datos o base de datos
	storedPassword := "$2a$10$NLwE9d4v7bSe1WdJAnD4Hu7xNDM1YsUkL3N7iWbS1t5BE8e8SOXm2" // Hash bcrypt de "password"

	// Verificar contraseña
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Si las credenciales son válidas, generar token JWT
	tokenString, err := GenerateJWT(username, "user") // Cambiar "user" por el rol del usuario obtenido desde la base de datos
	if err != nil {
		return "", errors.New("failed to generate JWT token")
	}

	return tokenString, nil
}

// GenerateJWT genera un token JWT válido para el usuario
func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas

	// Definir los claims del token
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your_issuer",
		},
	}

	// Crear token JWT usando los claims y la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken verifica la validez y la firma de un token JWT
func (s Service) ValidateToken(tokenString string) error {
	// Parsear el token con las claims definidas
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return err
	}

	// Verificar si el token es válido
	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
