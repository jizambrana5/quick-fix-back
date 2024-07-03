package user

import "gorm.io/gorm"

// Repository UserRepository implements the Storage interface for users
type Repository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}
