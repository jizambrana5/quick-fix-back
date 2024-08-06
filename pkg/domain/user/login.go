package user

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	catalog "github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) LoginUser(ctx context.Context, email string, password string) (domain.User, string, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, "", err
	}

	// Verificar la contraseña
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, "", catalog.InvalidUserCredentials
	}

	// Crear sesión
	token, _ := uuid.NewV4()
	session := entities.Session{
		UserID:    user.ID,
		Token:     token.String(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour), // Sesión expira en 24 horas
	}

	if err := s.storage.CreateSession(ctx, session); err != nil {
		return domain.User{}, "", err
	}

	return user, token.String(), nil
}

func (s *Service) LoginProfessional(ctx context.Context, email, password string) (domain.Professional, string, error) {
	professional, err := s.storage.GetProfessionalByEmail(ctx, email)
	if err != nil {
		return domain.Professional{}, "", err
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(professional.Password), []byte(password)); err != nil {
		return domain.Professional{}, "", catalog.InvalidProfessionalCredentials
	}

	// Crear sesión
	token, _ := uuid.NewV4()
	session := entities.Session{
		UserID:    professional.ID,
		Token:     token.String(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour), // Sesión expira en 24 horas
	}

	if err = s.storage.CreateSession(ctx, session); err != nil {
		return domain.Professional{}, "", err
	}

	return professional, token.String(), nil
}

func (s *Service) CreateSession(ctx context.Context, userID uint64) (string, error) {
	token, _ := uuid.NewV4()
	session := entities.Session{
		UserID:    userID,
		Token:     token.String(),
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	err := s.storage.CreateSession(ctx, session)
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (s *Service) ValidateSession(ctx context.Context, token string) (entities.Session, error) {
	return s.storage.GetSessionByToken(ctx, token)
}

func (s *Service) DeleteSession(ctx context.Context, token string) error {
	return s.storage.DeleteSession(ctx, token)
}

func (s *Service) DeleteExpiredSessions(ctx context.Context) error {
	return s.storage.DeleteExpiredSessions(ctx)
}
