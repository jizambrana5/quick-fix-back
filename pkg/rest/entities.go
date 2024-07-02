package rest

import (
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

const layout = "2006-01-02 15:04"

type (
	CreateOrderRequest struct {
		UserID         uint64 `json:"user_id"`
		ProfessionalID uint64 `json:"professional_id"`
		ScheduleTo     string `json:"schedule_to"`
	}

	AdvanceOrderRequest struct {
		Status domain.Status `json:"status"`
	}
)

func (r AdvanceOrderRequest) Validate() error {
	if r.Status != domain.OrderStatusPending && r.Status != domain.OrderStatusAccepted && r.Status != domain.OrderStatusCompleted &&
		r.Status != domain.OrderStatusCancelled {
		return errors.ErrInvalidInput
	}
	return nil
}

func (co CreateOrderRequest) Validate() error {
	if co.UserID == 0 {
		return errors.ErrInvalidUserID
	}
	if co.ProfessionalID == 0 {
		return errors.ErrInvalidProfessionalID
	}
	if co.ScheduleTo == "" {
		return errors.ErrInvalidScheduleTo
	}
	parsedTime, err := time.Parse(layout, co.ScheduleTo)
	if err != nil {
		return errors.ErrInvalidScheduleTo
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	timeInLoc := parsedTime.In(loc)
	if timeInLoc.Before(time.Now().In(loc)) {
		return errors.ErrInvalidScheduleTo
	}
	return nil
}
