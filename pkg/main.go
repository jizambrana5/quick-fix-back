package main

import (
	"fmt"
	"os"

	"github.com/jizambrana5/quickfix-back/pkg/domain/order"
	"github.com/jizambrana5/quickfix-back/pkg/repository/database"
	database2 "github.com/jizambrana5/quickfix-back/pkg/repository/database/order"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

func main() {

	// Inicializar la conexión a la base de datos
	//database.InitDatabase()

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	//TODO: Delete this
	fmt.Printf("DBHost: %s, DBUser: %s, DBPass: %s, DBName: %s, DBPort: %s",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db := database.NewRepository(database.Config{
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBPort:     dbPort,
	})

	orderRepo := database2.NewOrderRepository(db)
	orderSrv := order.NewService(orderRepo)
	handler := rest.NewHandler(orderSrv)

	server := rest.Routes(handler)
	err := server.Run(fmt.Sprintf("%s%s", ":", "8080"))
	if err != nil {
		panic(err)
	}
}
