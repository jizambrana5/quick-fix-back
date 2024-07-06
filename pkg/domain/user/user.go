package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

func (s *Service) RegisterUser(ctx context.Context, userReq entities.RegisterUserRequest) (domain.User, error) {
	userRepo, err := s.storage.GetUserByUsername(ctx, userReq.Username)
	if err != nil {
		return domain.User{}, errors.UserGet
	}
	if !userRepo.IsEmpty() {
		return domain.User{}, errors.UserAlreadyExists
	}

	userRepo, err = s.storage.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		return domain.User{}, errors.UserGet
	}
	if !userRepo.IsEmpty() {
		return domain.User{}, errors.UserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{}
	user.Password = string(hashedPassword)
	user.Role = "user"
	user.Email = userReq.Email
	user.Username = userReq.Username

	createdUser, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, errors.ErrCreateUser
	}

	return createdUser, nil
}

func (s *Service) RegisterProfessional(ctx context.Context, professionalReq entities.RegisterProfessionalRequest) (domain.Professional, error) {
	profRepo, err := s.storage.GetProfessionalByUsername(ctx, professionalReq.Username)
	if err != nil {
		return domain.Professional{}, errors.ProfessionalGet
	}
	if !profRepo.IsEmpty() {
		return domain.Professional{}, errors.ProfessionalAlreadyExist
	}

	profRepo, err = s.storage.GetProfessionalByEmail(ctx, professionalReq.Email)
	if err != nil {
		return domain.Professional{}, errors.ProfessionalGet
	}
	if !profRepo.IsEmpty() {
		return domain.Professional{}, errors.ProfessionalAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(professionalReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Professional{}, errors.ErrCreateProfessional
	}

	prof := domain.Professional{}
	prof.Password = string(hashedPassword)
	prof.Role = "professional"
	prof.Description = professionalReq.Description
	prof.Profession = professionalReq.Profession
	prof.Email = professionalReq.Email
	prof.Username = professionalReq.Username
	prof.Location.Department = professionalReq.Location.Department
	prof.Location.District = professionalReq.Location.District

	createdProfessional, err := s.storage.CreateProfessional(ctx, prof)
	if err != nil {
		return domain.Professional{}, errors.ErrCreateProfessional
	}

	return createdProfessional, nil
}

func (s *Service) GetUser(ctx context.Context, ID uint64) (domain.User, error) {
	user, err := s.storage.GetUserByID(ctx, ID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *Service) GetProfessional(ctx context.Context, ID uint64) (domain.Professional, error) {
	professional, err := s.storage.GetProfessionalByID(ctx, ID)
	if err != nil {
		return domain.Professional{}, err
	}
	return professional, nil
}

func (s *Service) FindProfessionalsByLocation(ctx context.Context, department, district string) ([]domain.Professional, error) {
	professionals, err := s.storage.FindProfessionalsByLocation(ctx, department, district)
	if err != nil {
		return nil, errors.ProfessionalGet
	}
	return professionals, nil
}
