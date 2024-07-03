package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

// CreateUser creates a new user in the database
func (r *Repository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	repoUser := FromDomainToUser(user)
	if err := r.DB.WithContext(ctx).Create(&repoUser).Error; err != nil {
		return domain.User{}, err
	}
	return repoUser.ToDomain(), nil
}

// CreateProfessional creates a new professional in the database
func (r *Repository) CreateProfessional(ctx context.Context, professional domain.Professional) (domain.Professional, error) {
	repoProfessional := FromDomainToProf(professional)
	if err := r.DB.WithContext(ctx).Create(&repoProfessional).Error; err != nil {
		return domain.Professional{}, err
	}
	return repoProfessional.ToDomain(), nil
}

// GetUserByUsername retrieves a user by username
func (r *Repository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	var user UserRepo
	if err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

// GetProfessionalByUsername retrieves a professional by username
func (r *Repository) GetProfessionalByUsername(ctx context.Context, username string) (domain.Professional, error) {
	var professional ProfessionalRepo
	if err := r.DB.WithContext(ctx).Where("username = ?", username).First(&professional).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Professional{}, nil
		}
		return domain.Professional{}, err
	}
	return professional.ToDomain(), nil
}

// GetUserByID retrieves a user by ID
func (r *Repository) GetUserByID(ctx context.Context, ID uint64) (domain.User, error) {
	var user UserRepo
	if err := r.DB.WithContext(ctx).First(&user, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

// GetProfessionalByID retrieves a professional by ID
func (r *Repository) GetProfessionalByID(ctx context.Context, ID uint64) (domain.Professional, error) {
	var professional ProfessionalRepo
	if err := r.DB.WithContext(ctx).First(&professional, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Professional{}, nil
		}
		return domain.Professional{}, err
	}
	return professional.ToDomain(), nil
}

// GetUserByEmail retrieves a user by email
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user UserRepo
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

// GetProfessionalByEmail retrieves a professional by email
func (r *Repository) GetProfessionalByEmail(ctx context.Context, email string) (domain.Professional, error) {
	var professional ProfessionalRepo
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&professional).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Professional{}, nil
		}
		return domain.Professional{}, err
	}
	return professional.ToDomain(), nil
}

func (r *Repository) FindProfessionalsByLocation(ctx context.Context, department, district string) ([]domain.Professional, error) {
	var professionals []ProfessionalRepo
	if err := r.DB.WithContext(ctx).Where("location_department = ? AND location_district = ?", department, district).Find(&professionals).Error; err != nil {
		return nil, err
	}
	var domainProfessionals []domain.Professional
	for _, prof := range professionals {
		domainProfessionals = append(domainProfessionals, prof.ToDomain())
	}
	return domainProfessionals, nil
}
