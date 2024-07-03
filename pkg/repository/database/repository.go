package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/jizambrana5/quickfix-back/pkg/repository/database/order"
	"github.com/jizambrana5/quickfix-back/pkg/repository/database/user"
)

var DB *gorm.DB

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func NewRepository(config Config) *gorm.DB {
	// Construir la cadena de conexión DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	// Abrir la conexión a PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database", err)
		panic(err)
	}

	// AutoMigrate creará las tablas, columnas faltantes e índices faltantes
	if err = db.AutoMigrate(&user.UserRepo{}, &user.ProfessionalRepo{}, &order.OrderRepo{}); err != nil {
		log.Fatalf("Error auto migrating models: %s", err)
		panic(err)
	}

	// Retornar una instancia de OrderRepository con la conexión establecida
	return db
}
