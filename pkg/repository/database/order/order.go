package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/repository/database"
)

// OrderRepository implements the Storage interface
type OrderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// GetOrder retrieves an order by ID
func (r *OrderRepository) GetOrder(ctx context.Context, ID string) (domain.Order, error) {
	var order database.OrderRepo
	if err := r.DB.WithContext(ctx).First(&order, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Order{}, nil
		}
		return domain.Order{}, err
	}
	return order.ToDomain(), nil
}

// CreateOrder creates a new order
func (r *OrderRepository) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	repoOrder := database.FromDomain(order)
	if err := r.DB.WithContext(ctx).Create(&repoOrder).Error; err != nil {
		return domain.Order{}, err
	}
	return repoOrder.ToDomain(), nil
}

// UpdateOrder updates an existing order
func (r *OrderRepository) UpdateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	repoOrder := database.FromDomain(order)
	if err := r.DB.WithContext(ctx).Save(&repoOrder).Error; err != nil {
		return domain.Order{}, err
	}
	return repoOrder.ToDomain(), nil
}

// FindOrdersByUserID finds orders by user ID
func (r *OrderRepository) FindOrdersByUserID(ctx context.Context, userID uint64) ([]domain.Order, error) {
	var orders []database.OrderRepo
	if err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}

// FindOrdersByProfessionalID finds orders by professional ID
func (r *OrderRepository) FindOrdersByProfessionalID(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
	var orders []database.OrderRepo
	if err := r.DB.WithContext(ctx).Where("professional_id = ?", professionalID).Find(&orders).Error; err != nil {
		return nil, err
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}

// FindOrdersByStatus finds orders by status
func (r *OrderRepository) FindOrdersByStatus(ctx context.Context, status string) ([]domain.Order, error) {
	var orders []database.OrderRepo
	if err := r.DB.WithContext(ctx).Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}

// FindOrdersBySchedule finds orders by schedule to, user id and professional id
func (r *OrderRepository) FindOrdersBySchedule(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
	var orders []database.OrderRepo
	result := r.DB.WithContext(ctx).Where("schedule_to = ? AND user_id = ? AND professional_id = ?", scheduleTo, userID, professionalID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}
