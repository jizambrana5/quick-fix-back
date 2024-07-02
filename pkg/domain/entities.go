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
)

type (
	Order struct {
		ID             string `json:"id"`
		UserID         uint64 `json:"user_id"`
		ProfessionalID uint64 `json:"professional_id"`
		Status         Status `json:"status"`
		Dates          Dates  `json:"dates"`
	}
	Dates struct {
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		ScheduleTo time.Time `json:"schedule_to"`
	}
	Status string
)

func (o Order) IsEmpty() bool {
	return o == Order{}
}

func (o Order) Validate() error {
	if o.Status == OrderStatusCompleted {
		return errors.OrderCompleted
	}

	if o.Status == OrderStatusCancelled {
		return errors.OrderCanceled
	}
	return nil
}
