package domain

import (
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

// OrderStatus represents the possible statuses of an order
const (
	OrderStatusPending   Status = "pending"
	OrderStatusAccepted  Status = "accepted"
	OrderStatusCompleted Status = "completed"
	OrderStatusCancelled Status = "cancelled"
	OrderStatusCreated   Status = "created"

	Plomero      Profession = "PLOMERO"
	Gasista      Profession = "GASISTA"
	Electricista Profession = "ELECTRICISTA"
)

type (
	Order struct {
		ID             string   `json:"id"`
		UserID         uint64   `json:"user_id"`
		ProfessionalID uint64   `json:"professional_id"`
		Status         Status   `json:"status"`
		Dates          Dates    `json:"dates"`
		Address        string   `json:"address"`
		Location       Location `json:"location"`
	}
	Dates struct {
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		ScheduleTo time.Time `json:"schedule_to"`
	}
	Status string

	User struct {
		ID        uint64    `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Password  string    `json:"-"`
		Name      string    `json:"name"`
		LastName  string    `json:"last_name"`
		Phone     string    `json:"phone"`
		Address   string    `json:"address"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Professional struct {
		ID          uint64     `json:"id"`
		Username    string     `json:"username"`
		Email       string     `json:"email"`
		Password    string     `json:"-"`
		Role        string     `json:"role"`
		Profession  Profession `json:"profession"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		Location    Location   `json:"location"`
		Name        string     `json:"name"`
		LastName    string     `json:"last_name"`
		Phone       string     `json:"phone"`
		Address     string     `json:"address"`
	}
	Location struct {
		Department string `json:"department"`
		District   string `json:"district"`
	}
	Profession string
)

// IsEmpty checks if the User struct is empty
func (u User) IsEmpty() bool {
	return u.ID == 0 && u.Username == "" && u.Email == "" && u.Password == "" && u.Role == ""
}

// IsEmpty checks if the Professional struct is empty
func (p Professional) IsEmpty() bool {
	return p.ID == 0 && p.Username == "" && p.Email == "" && p.Password == "" && p.Role == "" && p.Profession == "" &&
		p.Description == ""
}

func (o Order) IsEmpty() bool {
	return o == Order{}
}

func (o Order) ValidateToStatus(status Status) error {
	if o.Status == status {
		return errors.OrderAlreadyInRequestedStatus
	}

	if o.Status == OrderStatusCompleted {
		return errors.OrderCompleted
	}

	if o.Status == OrderStatusCancelled {
		return errors.OrderCanceled
	}
	return nil
}
