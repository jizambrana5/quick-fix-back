package order

import (
	"context"
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

func (s Service) GetOrder(ctx context.Context, ID string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, ID)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (s Service) GetOrdersByUser(ctx context.Context, userID uint64) ([]domain.Order, error) {
	orders, err := s.storage.FindOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s Service) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	now := time.Now()
	order.Dates.CreatedAt = now
	order.Dates.UpdatedAt = now

	createdOrder, err := s.storage.CreateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, err
	}
	return createdOrder, nil
}

func (s Service) AdvanceOrder(ctx context.Context, ID string, status string) (domain.Order, error) {
	order, err := s.storage.GetOrder(ctx, ID)
	if err != nil {
		return domain.Order{}, err
	}

	order.Status = domain.Status(status)
	order.Dates.UpdatedAt = time.Now()

	updatedOrder, err := s.storage.UpdateOrder(ctx, order)
	if err != nil {
		return domain.Order{}, err
	}
	return updatedOrder, nil
}
