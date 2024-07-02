package order

import (
	"context"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

type Service struct {
	storage Storage
}

type Storage interface {
	GetOrder(ctx context.Context, ID string) (domain.Order, error)
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	UpdateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	FindOrdersByUserID(ctx context.Context, userID uint64) ([]domain.Order, error)
	FindOrdersByProfessionalID(ctx context.Context, professionalID uint64) ([]domain.Order, error)
	FindOrdersByStatus(ctx context.Context, status string) ([]domain.Order, error)
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

var _ rest.OrderService = (*Service)(nil)
