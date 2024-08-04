package order

import (
	"context"
	"errors"
	"time"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	catalog "github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

func (t *orderSuite) Test_GetOrder_StorageError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	h, err := t.service.GetOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderGet)
	t.Empty(h)
}

func (t *orderSuite) Test_GetOrder_NotFoundError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{}, nil
	}
	h, err := t.service.GetOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderNotFound)
	t.Empty(h)
}

func (t *orderSuite) Test_GetOrder_OK() {
	h, err := t.service.GetOrder(t.ctx, "aaaa")
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *orderSuite) Test_GetOrdersByUser_StorageError() {
	t.storageMock.FindOrdersByUserIDFunc = func(ctx context.Context, userID uint64) ([]domain.Order, error) {
		return nil, errors.New("sth went wrong")
	}
	h, err := t.service.GetOrdersByUser(t.ctx, 1111)
	t.ErrorIs(err, catalog.OrdersGet)
	t.Empty(h)
}

func (t *orderSuite) Test_GetOrdersByUser_Ok() {
	h, err := t.service.GetOrdersByUser(t.ctx, 1111)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *orderSuite) Test_GetOrdersByProfessional_StorageError() {
	t.storageMock.FindOrdersByProfessionalIDFunc = func(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
		return nil, errors.New("sth went wrong")
	}
	h, err := t.service.GetOrdersByProfessional(t.ctx, 2222)
	t.ErrorIs(err, catalog.OrdersGet)
	t.Empty(h)
}

func (t *orderSuite) Test_GetOrdersByProfessional_Ok() {
	h, err := t.service.GetOrdersByProfessional(t.ctx, 2222)
	t.Nil(err)
	t.NotEmpty(h)
}

func (t *orderSuite) Test_CreateOrder_InvalidScheduleTo() {
	_, err := t.service.CreateOrder(t.ctx, entities.CreateOrderRequest{ScheduleTo: "invalid"})
	t.ErrorIs(err, catalog.ErrInvalidScheduleTo)
	t.Len(t.storageMock.FindOrdersByScheduleCalls(), 0)
}

func (t *orderSuite) Test_CreateOrder_OrdersGetError() {
	t.storageMock.FindOrdersByScheduleFunc = func(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
		return nil, errors.New("sth went wrong")
	}
	_, err := t.service.CreateOrder(t.ctx, entities.CreateOrderRequest{ScheduleTo: time.Now().AddDate(0, 0, 2).Format(layout)})
	t.ErrorIs(err, catalog.OrdersGet)
	t.Len(t.storageMock.FindOrdersByScheduleCalls(), 1)
}

func (t *orderSuite) Test_CreateOrder_OrderAlreadyExists() {
	t.storageMock.FindOrdersByScheduleFunc = func(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
		return []domain.Order{{}}, nil
	}
	_, err := t.service.CreateOrder(t.ctx, entities.CreateOrderRequest{ScheduleTo: time.Now().AddDate(0, 0, 2).Format(layout)})
	t.Nil(err)
	t.Len(t.storageMock.FindOrdersByScheduleCalls(), 1)
}

func (t *orderSuite) Test_CreateOrder_OK() {
	t.storageMock.FindOrdersByScheduleFunc = func(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
		return []domain.Order{}, nil
	}
	_, err := t.service.CreateOrder(t.ctx, entities.CreateOrderRequest{ScheduleTo: time.Now().AddDate(0, 0, 2).Format(layout)})
	t.Nil(err)
	t.Len(t.storageMock.FindOrdersByScheduleCalls(), 1)
	t.Len(t.storageMock.CreateOrderCalls(), 1)
}

func (t *orderSuite) Test_AcceptOrder_StorageError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.AcceptOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderGet)
	t.Empty(order)
}

func (t *orderSuite) Test_AcceptOrder_OrderAlreadyInRequestedStatus() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusAccepted}, nil
	}
	order, err := t.service.AcceptOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderAlreadyInRequestedStatus)
	t.Empty(order)
}

func (t *orderSuite) Test_AcceptOrder_UpdateOrderError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.AcceptOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderUpdate)
	t.Empty(order)
}

func (t *orderSuite) Test_AcceptOrder_OK() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return order, nil
	}
	order, err := t.service.AcceptOrder(t.ctx, "aaaa")
	t.Nil(err)
	t.NotEmpty(order)
}

func (t *orderSuite) Test_CompleteOrder_GetOrderError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.CompleteOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderGet)
	t.Empty(order)
}

func (t *orderSuite) Test_CompleteOrder_OrderAlreadyInRequestedStatus() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusCompleted}, nil
	}
	order, err := t.service.CompleteOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderAlreadyInRequestedStatus)
	t.Empty(order)
}

func (t *orderSuite) Test_CompleteOrder_UpdateOrderError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.CompleteOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderUpdate)
	t.Empty(order)
}

func (t *orderSuite) Test_CompleteOrder_OK() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return order, nil
	}
	order, err := t.service.CompleteOrder(t.ctx, "aaaa")
	t.Nil(err)
	t.NotEmpty(order)
}

func (t *orderSuite) Test_CancelOrder_GetOrderError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.CancelOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderGet)
	t.Empty(order)
}

func (t *orderSuite) Test_CancelOrder_OrderAlreadyInRequestedStatus() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusCancelled}, nil
	}
	order, err := t.service.CancelOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderAlreadyInRequestedStatus)
	t.Empty(order)
}

func (t *orderSuite) Test_CancelOrder_UpdateOrderError() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return domain.Order{}, errors.New("sth went wrong")
	}
	order, err := t.service.CancelOrder(t.ctx, "aaaa")
	t.ErrorIs(err, catalog.OrderUpdate)
	t.Empty(order)
}

func (t *orderSuite) Test_CancelOrder_OK() {
	t.storageMock.GetOrderFunc = func(ctx context.Context, ID string) (domain.Order, error) {
		return domain.Order{Status: domain.OrderStatusPending}, nil
	}
	t.storageMock.UpdateOrderFunc = func(ctx context.Context, order domain.Order) (domain.Order, error) {
		return order, nil
	}
	order, err := t.service.CancelOrder(t.ctx, "aaaa")
	t.Nil(err)
	t.NotEmpty(order)
}
