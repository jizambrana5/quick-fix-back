package database

import (
	"errors"

	"gorm.io/gorm"
)

// GetUserByID obtiene un usuario por su ID
func GetUserByID(id uint) (*User, error) {
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No se encontró el usuario
		}
		return nil, err // Otro error
	}
	return &user, nil
}

// CreateUser crea un nuevo usuario
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// UpdateUser actualiza los datos de un usuario
func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

// DeleteUser elimina un usuario por su ID
func DeleteUser(id uint) error {
	return DB.Delete(&User{}, id).Error
}
