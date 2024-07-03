package user

import (
	"context"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

type Service struct {
	storage Storage
}

type Storage interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	CreateProfessional(ctx context.Context, professional domain.Professional) (domain.Professional, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	GetProfessionalByUsername(ctx context.Context, username string) (domain.Professional, error)
	GetUserByID(ctx context.Context, ID uint64) (domain.User, error)
	GetProfessionalByID(ctx context.Context, ID uint64) (domain.Professional, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetProfessionalByEmail(ctx context.Context, email string) (domain.Professional, error)
	FindProfessionalsByLocation(ctx context.Context, department, district string) ([]domain.Professional, error)
}

func NewUserService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

var _ rest.UserService = (*Service)(nil)
