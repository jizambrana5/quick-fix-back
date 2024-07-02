package domain

import "time"

// OrderStatus represents the possible statuses of an order
const (
	OrderStatusPending   Status = "pending"
	OrderStatusAccepted  Status = "accepted"
	OrderStatusCompleted Status = "completed"
	OrderStatusCancelled Status = "cancelled"
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
