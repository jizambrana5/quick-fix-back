package user

import (
	"context"
	"github.com/jizambrana5/quickfix-back/pkg/domain"
	catalog "github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) LoginUser(ctx context.Context, email string, password string) (domain.User, error) {
	user, err := s.storage.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, catalog.ErrInvalidEmail
	}

	// Verificar la contraseña
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.User{}, catalog.InvalidUserCredentials
	}

	return user, nil
}

func (s *Service) LoginProfessional(ctx context.Context, email, password string) (domain.Professional, error) {
	professional, err := s.storage.GetProfessionalByEmail(ctx, email)
	if err != nil {
		return domain.Professional{}, catalog.ErrInvalidEmail
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(professional.Password), []byte(password)); err != nil {
		return domain.Professional{}, catalog.InvalidProfessionalCredentials
	}

	return professional, nil
}
