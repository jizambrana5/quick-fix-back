package database

// User estructura para los datos del usuario
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role
}

// Role estructura para los roles de usuario
type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
