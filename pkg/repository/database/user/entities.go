package user

import (
	"github.com/jizambrana5/quickfix-back/pkg/domain"

	"time"
)

// UserRepo represents the user entity in the database
type UserRepo struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ProfessionalRepo represents the professional entity in the database
type ProfessionalRepo struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement"`
	Username    string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email       string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string       `gorm:"not null"`
	Role        string       `gorm:"type:varchar(50);not null"`
	Profession  string       `gorm:"type:varchar(100);not null"`
	Description string       `gorm:"type:text"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	Location    LocationRepo `gorm:"embedded;embeddedPrefix:location_"` // Embedded struct
}

// LocationRepo es la entidad del repositorio que representa la ubicación de un profesional.
type LocationRepo struct {
	Department string `gorm:"type:varchar(100);not null"`
	District   string `gorm:"type:varchar(100);not null"`
}

// ToDomain converts UserRepo to domain.User
func (u *UserRepo) ToDomain() domain.User {
	return domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromDomainToUser converts domain.User to UserRepo
func FromDomainToUser(u domain.User) UserRepo {
	return UserRepo{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ToDomain converts ProfessionalRepo to domain.Professional
func (p *ProfessionalRepo) ToDomain() domain.Professional {
	return domain.Professional{
		ID:          p.ID,
		Username:    p.Username,
		Email:       p.Email,
		Role:        p.Role,
		Profession:  p.Profession,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Location: domain.Location{
			Department: p.Location.Department,
			District:   p.Location.District,
		},
	}
}

// FromDomainToProf converts domain.Professional to ProfessionalRepo
func FromDomainToProf(p domain.Professional) ProfessionalRepo {
	return ProfessionalRepo{
		ID:          p.ID,
		Username:    p.Username,
		Email:       p.Email,
		Password:    p.Password,
		Role:        p.Role,
		Profession:  p.Profession,
		Description: p.Description,
		Location: LocationRepo{
			Department: p.Location.Department,
			District:   p.Location.District,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
