package order

import "gorm.io/gorm"

// Repository OrderRepository implements the Storage interface
type Repository struct {
	DB *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
