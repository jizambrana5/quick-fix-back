package order

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
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
	orders, err := s.storage.FindOrdersByProfessionalID(ctx, professionalID)
	if err != nil {
		return nil, errors.OrdersGet
	}
	return orders, nil
}

func (s Service) CreateOrder(ctx context.Context, orderReq entities.CreateOrderRequest) (domain.Order, error) {
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
		for _, order := range orders {
			if order.Status == domain.OrderStatusAccepted {
				return domain.Order{}, errors.OrderAlreadyExists
			}
			if order.Status == domain.OrderStatusCompleted {
				return domain.Order{}, errors.OrderCompleted
			}
		}
	}

	orderID, _ := uuid.NewV4()

	order := domain.Order{
		ID:             orderID.String(),
		UserID:         orderReq.UserID,
		ProfessionalID: orderReq.ProfessionalID,
		Dates: domain.Dates{
			CreatedAt:  time.Now().In(loc),
			UpdatedAt:  time.Now().In(loc),
			ScheduleTo: parsedTime,
		},
		Status:  domain.OrderStatusPending,
		Address: orderReq.Address,
		Location: domain.Location{
			Department: orderReq.Location.Department,
			District:   orderReq.Location.District,
		},
		Description: orderReq.Description,
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

	if err = order.ValidateToStatus(domain.OrderStatusAccepted); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusAccepted
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	order.Dates.UpdatedAt = time.Now().In(loc)

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

	if err = order.ValidateToStatus(domain.OrderStatusCompleted); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusCompleted
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	order.Dates.UpdatedAt = time.Now().In(loc)

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

	if err = order.ValidateToStatus(domain.OrderStatusCancelled); err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.OrderStatusCancelled
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	order.Dates.UpdatedAt = time.Now().In(loc)

	updatedOrder, err := s.storage.UpdateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, errors.OrderUpdate
	}
	return updatedOrder, nil
}

func (s Service) GetOrdersByProfessionalAndScheduleTo(ctx context.Context, professionalID uint64, scheduleTo time.Time) ([]domain.Order, error) {
	orders, err := s.storage.FindOrdersByProfessionalAndBySchedule(ctx, professionalID, scheduleTo)
	if err != nil {
		return nil, errors.OrdersGet
	}
	return orders, nil
}

func (s Service) GetOrdersByProfessionalAndDay(ctx context.Context, id uint64, day time.Time) ([]domain.Order, error) {
	orders, err := s.storage.FindOrdersByProfessionalAndByDay(ctx, id, day)
	if err != nil {
		return nil, errors.OrdersGet
	}
	return orders, nil
}
