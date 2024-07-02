package database

import (
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

// OrderRepo represents the order entity in the database
type OrderRepo struct {
	ID             string    `gorm:"primaryKey;autoIncrement"`
	UserID         uint64    `gorm:"not null"`
	ProfessionalID uint64    `gorm:"not null"`
	Status         string    `gorm:"type:varchar(20);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	ScheduleTo     time.Time `gorm:"autoScheduleTime"`
}

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
	}
}
