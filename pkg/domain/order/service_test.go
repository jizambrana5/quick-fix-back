package order

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/domain/order/mocks"
)

type orderSuite struct {
	suite.Suite
	ctx         context.Context
	storageMock *mocks.StorageMock
	service     *Service
}

func (t *orderSuite) SetupTest() {
	t.ctx = context.Background()
	t.storageMock = &mocks.StorageMock{
		CreateOrderFunc: func(ctx context.Context, order domain.Order) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		FindOrdersByProfessionalIDFunc: func(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaaa"}}, nil
		},
		FindOrdersByScheduleFunc: func(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaaa"}}, nil
		},
		FindOrdersByStatusFunc: func(ctx context.Context, status string) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaaa"}}, nil
		},
		FindOrdersByUserIDFunc: func(ctx context.Context, userID uint64) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaaa"}}, nil
		},
		GetOrderFunc: func(ctx context.Context, ID string) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		UpdateOrderFunc: func(ctx context.Context, order domain.Order) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
	}
	t.service = NewService(t.storageMock)
}

func (t *orderSuite) Test_NewService() {
	t.NotNil(NewService(t.storageMock))
}

func TestOrder(t *testing.T) {
	suite.Run(t, new(orderSuite))
}
