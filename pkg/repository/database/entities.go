package database

import "gorm.io/gorm"

// User representa un usuario en la base de datos
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	// Otros campos según tus necesidades
}
