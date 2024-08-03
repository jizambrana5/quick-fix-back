package order

import (
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/repository/database/user"
)

// OrderRepo represents the order entity in the database
type (
	OrderRepo struct {
		ID             string                `gorm:"primaryKey;autoIncrement"`
		UserID         uint64                `gorm:"not null"`
		User           user.UserRepo         `gorm:"foreignKey:UserID"`
		ProfessionalID uint64                `gorm:"not null"`
		Professional   user.ProfessionalRepo `gorm:"foreignKey:ProfessionalID"`
		Status         string                `gorm:"type:varchar(20);not null"`
		CreatedAt      time.Time             `gorm:"autoCreateTime"`
		UpdatedAt      time.Time             `gorm:"autoUpdateTime"`
		ScheduleTo     time.Time             `gorm:"autoScheduleTime"`
		Address        string                `gorm:"size:255;not null"`
		Location       LocationRepo          `gorm:"embedded;embeddedPrefix:location_"` // Embedded struct
		Description    string                `gorm:"type:text;not null"`
	}
	// LocationRepo es la entidad del repositorio que representa la ubicación de un profesional.
	LocationRepo struct {
		Department string `gorm:"type:varchar(100);not null"`
		District   string `gorm:"type:varchar(100);not null"`
	}
)

// ToDomain transforms the repository order entity to the domain order entity
func (o *OrderRepo) ToDomain() domain.Order {
	return domain.Order{
		ID:             o.ID,
		UserID:         o.UserID,
		ProfessionalID: o.ProfessionalID,
		Status:         domain.Status(o.Status),
		Dates: domain.Dates{
			CreatedAt:  o.CreatedAt,
			UpdatedAt:  o.UpdatedAt,
			ScheduleTo: o.ScheduleTo,
		},
		Address: o.Address,
		Location: domain.Location{
			Department: o.Location.Department,
			District:   o.Location.District,
		},
		Description: o.Description,
	}
}

// FromDomain transforms the domain order entity to the repository order entity
func FromDomain(order domain.Order) OrderRepo {
	return OrderRepo{
		ID:             order.ID,
		UserID:         order.UserID,
		ProfessionalID: order.ProfessionalID,
		Status:         string(order.Status),
		CreatedAt:      order.Dates.CreatedAt,
		UpdatedAt:      order.Dates.UpdatedAt,
		ScheduleTo:     order.Dates.ScheduleTo,
		Address:        order.Address,
		Location: LocationRepo{
			Department: order.Location.Department,
			District:   order.Location.District,
		},
		Description: order.Description,
	}
}
