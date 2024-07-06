package user

import (
	"context"
	"errors"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
)

func (t *userSuite) Test_RegisterUser_GetUserByNameError() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, errors.New("sth went wrong")
	}

	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterUser_UsernameAlreadyUsedError() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{
			Username: "username",
		}, nil
	}

	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterUser_GetByEmailError() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, nil
	}
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, errors.New("sth went wrong")
	}

	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterUser_EmailAlreadyUsedError() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, nil
	}
	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterUser_CreateUserError() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, nil
	}
	t.storageMock.GetUserByEmailFunc = func(ctx context.Context, email string) (domain.User, error) {
		return domain.User{}, nil
	}
	t.storageMock.CreateUserFunc = func(ctx context.Context, user domain.User) (domain.User, error) {
		return domain.User{}, errors.New("sth went wrong")
	}
	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterUser_OK() {
	t.storageMock.GetUserByUsernameFunc = func(ctx context.Context, username string) (domain.User, error) {
		return domain.User{}, nil
	}
	t.storageMock.GetUserByEmailFunc = func(ctx context.Context, email string) (domain.User, error) {
		return domain.User{}, nil
	}
	h, err := t.service.RegisterUser(t.ctx, entities.RegisterUserRequest{
		Username: "username",
		Email:    "asd@asd.com",
		Password: "password",
	})
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *userSuite) Test_RegisterProfessional_GetProfessionalByNameError() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{}, errors.New("sth went wrong")
	}

	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterProfessional_UsernameAlreadyUsedError() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{Username: "username"}, nil
	}

	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterProfessional_GetByEmailError() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	t.storageMock.GetProfessionalByEmailFunc = func(ctx context.Context, email string) (domain.Professional, error) {
		return domain.Professional{}, errors.New("sth went wrong")
	}

	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterProfessional_EmailAlreadyUsedError() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	t.storageMock.GetProfessionalByEmailFunc = func(ctx context.Context, email string) (domain.Professional, error) {
		return domain.Professional{Username: "username"}, nil
	}

	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterProfessional_CreateUserError() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	t.storageMock.GetProfessionalByEmailFunc = func(ctx context.Context, email string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	t.storageMock.CreateProfessionalFunc = func(ctx context.Context, professional domain.Professional) (domain.Professional, error) {
		return domain.Professional{}, errors.New("sth went wrong")
	}
	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_RegisterProfessional_OK() {
	t.storageMock.GetProfessionalByUsernameFunc = func(ctx context.Context, username string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	t.storageMock.GetProfessionalByEmailFunc = func(ctx context.Context, email string) (domain.Professional, error) {
		return domain.Professional{}, nil
	}
	h, err := t.service.RegisterProfessional(t.ctx, entities.RegisterProfessionalRequest{
		Username: "username",
	})
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *userSuite) Test_GetUser_GetUserError() {
	t.storageMock.GetUserByIDFunc = func(ctx context.Context, ID uint64) (domain.User, error) {
		return domain.User{}, errors.New("sth went wrong")
	}

	h, err := t.service.GetUser(t.ctx, 1111)
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_GetUser_Ok() {
	h, err := t.service.GetUser(t.ctx, 1111)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *userSuite) Test_GetProfessional_GetProfessionalError() {
	t.storageMock.GetProfessionalByIDFunc = func(ctx context.Context, ID uint64) (domain.Professional, error) {
		return domain.Professional{}, errors.New(" sth went wrong")
	}

	h, err := t.service.GetProfessional(t.ctx, 2222)
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_GetProfessional_Ok() {
	h, err := t.service.GetProfessional(t.ctx, 2222)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *userSuite) Test_FindProfessionalByLocation_StorageError() {
	t.storageMock.FindProfessionalsByLocationFunc = func(ctx context.Context, department string, district string) ([]domain.Professional, error) {
		return nil, errors.New("sth went wrong")
	}
	h, err := t.service.FindProfessionalsByLocation(t.ctx, "department", "district")
	t.NotNil(err)
	t.Empty(h)
}

func (t *userSuite) Test_FindProfessionalByLocation_Ok() {
	h, err := t.service.FindProfessionalsByLocation(t.ctx, "department", "district")
	t.Nil(err)
	t.NotEmpty(h)
}
