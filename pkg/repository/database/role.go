package database

import (
	"github.com/jinzhu/gorm"
)

// UserRepository define métodos para interactuar con la tabla de usuarios
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository inicializa un nuevo repositorio de usuarios
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FindByUsername busca un usuario por nombre de usuario
func (r *UserRepository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// RoleRepository define métodos para interactuar con la tabla de roles
type RoleRepository struct {
	DB *gorm.DB
}

// NewRoleRepository inicializa un nuevo repositorio de roles
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

// FindByName busca un rol por nombre
func (r *RoleRepository) FindByName(name string) (*Role, error) {
	var role Role
	if err := r.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
