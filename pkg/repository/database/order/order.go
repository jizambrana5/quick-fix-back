package database

import (
	"context"

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
	if err := r.DB.WithContext(ctx).Preload("User").Preload("Professional").First(&order, "id = ?", ID).Error; err != nil {
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
	if err := r.DB.WithContext(ctx).Preload("User").Preload("Professional").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
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
	if err := r.DB.WithContext(ctx).Preload("User").Preload("Professional").Where("professional_id = ?", professionalID).Find(&orders).Error; err != nil {
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
	if err := r.DB.WithContext(ctx).Preload("User").Preload("Professional").Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}
