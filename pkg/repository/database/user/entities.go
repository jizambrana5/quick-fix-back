package user

import (
	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"gorm.io/gorm"
	"time"
)

// UserRepo represents the user entity in the database
type UserRepo struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Address   string    `gorm:"type:varchar(100);not null"`
	Phone     string    `gorm:"type:varchar(100);not null"`
	Role      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ProfessionalRepo represents the professional entity in the database
type ProfessionalRepo struct {
	ID                 uint64       `gorm:"primaryKey;autoIncrement"`
	Username           string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email              string       `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password           string       `gorm:"not null"`
	Name               string       `gorm:"type:varchar(100);not null"`
	LastName           string       `gorm:"type:varchar(100);not null"`
	Address            string       `gorm:"type:varchar(100);not null"`
	Phone              string       `gorm:"type:varchar(100);not null"`
	Role               string       `gorm:"type:varchar(50);not null"`
	Profession         string       `gorm:"type:varchar(100);not null"`
	Description        string       `gorm:"type:text"`
	CreatedAt          time.Time    `gorm:"autoCreateTime"`
	UpdatedAt          time.Time    `gorm:"autoUpdateTime"`
	Location           LocationRepo `gorm:"embedded;embeddedPrefix:location_"` // Embedded struct
	RegistrationNumber string       `gorm:"type:varchar(50);not null"`
}

// LocationRepo es la entidad del repositorio que representa la ubicación de un profesional.
type LocationRepo struct {
	Department string `gorm:"type:varchar(100);not null"`
	District   string `gorm:"type:varchar(100);not null"`
}

type SessionRepo struct {
	SessionID int64     `gorm:"primaryKey" json:"session_id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"not null" json:"token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

// ToDomain converts UserRepo to domain.User
func (u *UserRepo) ToDomain() domain.User {
	return domain.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		Name:      u.Name,
		LastName:  u.LastName,
		Address:   u.Address,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Password:  u.Password,
	}
}

// FromDomainToUser converts domain.User to UserRepo
func FromDomainToUser(u domain.User) UserRepo {
	return UserRepo{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Name:      u.Name,
		LastName:  u.LastName,
		Address:   u.Address,
		Phone:     u.Phone,
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
		Profession:  domain.Profession(p.Profession),
		Description: p.Description,
		Name:        p.Name,
		LastName:    p.LastName,
		Address:     p.Address,
		Phone:       p.Phone,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Location: domain.Location{
			Department: p.Location.Department,
			District:   p.Location.District,
		},
		RegistrationNumber: p.RegistrationNumber,
		Password:           p.Password,
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
		Profession:  string(p.Profession),
		Description: p.Description,
		Location: LocationRepo{
			Department: p.Location.Department,
			District:   p.Location.District,
		},
		Name:               p.Name,
		LastName:           p.LastName,
		Phone:              p.Phone,
		Address:            p.Address,
		CreatedAt:          p.CreatedAt,
		UpdatedAt:          p.UpdatedAt,
		RegistrationNumber: p.RegistrationNumber,
	}
}

// BeforeCreate hook de Gorm para convertir CreatedAt a UTC-3 antes de crear el registro
func (u *UserRepo) BeforeCreate(tx *gorm.DB) (err error) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return err
	}
	u.CreatedAt = time.Now().In(loc)
	return
}

// BeforeUpdate hook de Gorm para convertir UpdatedAt a UTC-3 antes de actualizar el registro
func (u *UserRepo) BeforeUpdate(tx *gorm.DB) (err error) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return err
	}
	u.UpdatedAt = time.Now().In(loc)
	return
}

// BeforeCreate hook de Gorm para convertir CreatedAt a UTC-3 antes de crear el registro
func (p *ProfessionalRepo) BeforeCreate(tx *gorm.DB) (err error) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return err
	}
	p.CreatedAt = time.Now().In(loc)
	return
}

// BeforeUpdate hook de Gorm para convertir UpdatedAt a UTC-3 antes de actualizar el registro
func (p *ProfessionalRepo) BeforeUpdate(tx *gorm.DB) (err error) {
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return err
	}
	p.UpdatedAt = time.Now().In(loc)
	return
}

// FromDomainToSession convierte una entidad de dominio a una entidad del repositorio
func FromDomainToSession(session entities.Session) SessionRepo {
	return SessionRepo{
		SessionID: session.SessionID,
		UserID:    session.UserID,
		Token:     session.Token,
		CreatedAt: session.CreatedAt,
		ExpiresAt: session.ExpiresAt,
	}
}

// ToDomain convierte una entidad del repositorio a una entidad de dominio
func (s SessionRepo) ToDomain() entities.Session {
	return entities.Session{
		SessionID: s.SessionID,
		UserID:    s.UserID,
		Token:     s.Token,
		CreatedAt: s.CreatedAt,
		ExpiresAt: s.ExpiresAt,
	}
}
