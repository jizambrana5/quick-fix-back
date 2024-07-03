package order

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

const layout = "2006-01-02 15:04"

func (s Service) GetOrder(ctx context.Context, ID string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, ID)
	if err != nil {
		return domain.Order{}, errors.OrderGet
	}
	if order.IsEmpty() {
		return domain.Order{}, errors.OrderNotFound
	}
	return order, nil
}

func (s Service) GetOrdersByUser(ctx context.Context, userID uint64) ([]domain.Order, error) {
	orders, err := s.storage.FindOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errors.OrdersGet
	}
	return orders, nil
}

func (s Service) GetOrdersByProfessional(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
	orders, err := s.storage.FindOrdersByUserID(ctx, professionalID)
	if err != nil {
		return nil, errors.OrdersGet
	}
	return orders, nil
}

func (s Service) CreateOrder(ctx context.Context, orderReq rest.CreateOrderRequest) (domain.Order, error) {
	// Check if exist in that schedule to
	parsedTime, err := time.Parse(layout, orderReq.ScheduleTo)
	if err != nil {
		return domain.Order{}, errors.ErrInvalidScheduleTo
	}
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	timeInLoc := parsedTime.In(loc)

	orders, err := s.storage.FindOrdersBySchedule(ctx, timeInLoc, orderReq.UserID, orderReq.ProfessionalID)
	if err != nil {
		return domain.Order{}, errors.OrdersGet
	}

	if len(orders) > 0 {
		return domain.Order{}, errors.OrderAlreadyExists
	}

	now := time.Now()
	orderID, _ := uuid.NewV4()

	order := domain.Order{
		ID:             orderID.String(),
		UserID:         orderReq.UserID,
		ProfessionalID: orderReq.ProfessionalID,
		Dates: domain.Dates{
			CreatedAt:  now,
			UpdatedAt:  now,
			ScheduleTo: parsedTime,
		},
		Status: domain.OrderStatusPending,
	}

	createdOrder, err := s.storage.CreateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, errors.OrderSave
	}
	return createdOrder, nil
}

func (s Service) AcceptOrder(ctx context.Context, orderID string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, orderID)
	if err != nil {
		return domain.Order{}, errors.OrderGet
	}

	if order.IsEmpty() {
		return domain.Order{}, errors.OrderNotFound
	}

	if err = order.Validate(); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusAccepted
	order.Dates.UpdatedAt = time.Now()

	updatedOrder, err := s.storage.UpdateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, errors.OrderUpdate
	}
	return updatedOrder, nil
}

func (s Service) CompleteOrder(ctx context.Context, orderID string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, orderID)
	if err != nil {
		return domain.Order{}, errors.OrderGet
	}

	if order.IsEmpty() {
		return domain.Order{}, errors.OrderNotFound
	}

	if err = order.Validate(); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusCompleted
	order.Dates.UpdatedAt = time.Now()

	updatedOrder, err := s.storage.UpdateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, errors.OrderUpdate
	}
	return updatedOrder, nil
}

func (s Service) CancelOrder(ctx context.Context, orderID string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, orderID)
	if err != nil {
		return domain.Order{}, errors.OrderGet
	}

	if order.IsEmpty() {
		return domain.Order{}, errors.OrderNotFound
	}

	if err = order.Validate(); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusCancelled
	order.Dates.UpdatedAt = time.Now()

	updatedOrder, err := s.storage.UpdateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, errors.OrderUpdate
	}
	return updatedOrder, nil
}
