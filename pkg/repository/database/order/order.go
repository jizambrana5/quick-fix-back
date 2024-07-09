package order

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

// GetOrder retrieves an order by ID
func (r *Repository) GetOrder(ctx context.Context, ID string) (domain.Order, error) {
	var order OrderRepo
	if err := r.DB.WithContext(ctx).First(&order, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Order{}, nil
		}
		return domain.Order{}, err
	}
	return order.ToDomain(), nil
}

// CreateOrder creates a new order
func (r *Repository) CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	repoOrder := FromDomain(order)
	if err := r.DB.WithContext(ctx).Create(&repoOrder).Error; err != nil {
		return domain.Order{}, err
	}
	return repoOrder.ToDomain(), nil
}

// UpdateOrder updates an existing order
func (r *Repository) UpdateOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	repoOrder := FromDomain(order)
	if err := r.DB.WithContext(ctx).Save(&repoOrder).Error; err != nil {
		return domain.Order{}, err
	}
	return repoOrder.ToDomain(), nil
}

// FindOrdersByUserID finds orders by user ID
func (r *Repository) FindOrdersByUserID(ctx context.Context, userID uint64) ([]domain.Order, error) {
	var orders []OrderRepo
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
func (r *Repository) FindOrdersByProfessionalID(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
	var orders []OrderRepo
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
func (r *Repository) FindOrdersByStatus(ctx context.Context, status string) ([]domain.Order, error) {
	var orders []OrderRepo
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
func (r *Repository) FindOrdersBySchedule(ctx context.Context, scheduleTo time.Time, userID uint64, professionalID uint64) ([]domain.Order, error) {
	var orders []OrderRepo
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

// FindOrdersByProfessionalAndBySchedule finds orders by professional ID and schedule to
func (r *Repository) FindOrdersByProfessionalAndBySchedule(ctx context.Context, professionalID uint64, scheduleTo time.Time) ([]domain.Order, error) {
	var orders []OrderRepo
	result := r.DB.WithContext(ctx).Where("professional_id = ? AND schedule_to = ?", professionalID, scheduleTo).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}

func (r *Repository) FindOrdersByProfessionalAndByDay(ctx context.Context, professionalID uint64, day time.Time) ([]domain.Order, error) {
	// Hacer la consulta a la base de datos filtrando por professional ID y el día del schedule to si se proporciona
	db := r.DB.WithContext(ctx).Where("professional_id = ?", professionalID)
	if !day.IsZero() {
		db = db.Where("DATE(schedule_to) = ?", day.Format("2006-01-02"))
	}
	var orders []OrderRepo
	if err := db.Order("schedule_to ASC").Find(&orders).Error; err != nil {
		return nil, err
	}

	// Transformar a dominio
	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, order.ToDomain())
	}
	return domainOrders, nil
}
