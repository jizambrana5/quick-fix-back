package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/domain/user/mocks"
)

type userSuite struct {
	suite.Suite
	ctx         context.Context
	storageMock *mocks.StorageMock
	service     *Service
}

func (t *userSuite) SetupTest() {
	t.ctx = context.Background()
	t.storageMock = &mocks.StorageMock{
		CreateProfessionalFunc: func(ctx context.Context, professional domain.Professional) (domain.Professional, error) {
			return domain.Professional{ID: 1111}, nil
		},
		CreateUserFunc: func(ctx context.Context, user domain.User) (domain.User, error) {
			return domain.User{ID: 2222}, nil
		},
		FindProfessionalsByLocationFunc: func(ctx context.Context, department string, district string) ([]domain.Professional, error) {
			return []domain.Professional{{ID: 2222}}, nil
		},
		GetProfessionalByEmailFunc: func(ctx context.Context, email string) (domain.Professional, error) {
			return domain.Professional{ID: 1111}, nil
		},
		GetProfessionalByIDFunc: func(ctx context.Context, ID uint64) (domain.Professional, error) {
			return domain.Professional{ID: 1111}, nil
		},
		GetProfessionalByUsernameFunc: func(ctx context.Context, username string) (domain.Professional, error) {
			return domain.Professional{ID: 1111}, nil
		},
		GetUserByEmailFunc: func(ctx context.Context, email string) (domain.User, error) {
			return domain.User{ID: 2222}, nil
		},
		GetUserByIDFunc: func(ctx context.Context, ID uint64) (domain.User, error) {
			return domain.User{ID: 2222}, nil
		},
		GetUserByUsernameFunc: func(ctx context.Context, username string) (domain.User, error) {
			return domain.User{ID: 2222}, nil
		},
	}
	t.service = NewUserService(t.storageMock)
}

func (t *userSuite) Test_NewService() {
	t.NotNil(NewUserService(t.storageMock))
}

func TestUser(t *testing.T) {
	suite.Run(t, new(userSuite))
}
