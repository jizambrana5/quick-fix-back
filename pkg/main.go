package main

import (
	"fmt"
	"os"

	"github.com/jizambrana5/quickfix-back/pkg/domain/order"
	"github.com/jizambrana5/quickfix-back/pkg/domain/user"
	"github.com/jizambrana5/quickfix-back/pkg/lib/logs"
	"github.com/jizambrana5/quickfix-back/pkg/repository/database"
	orderRepo "github.com/jizambrana5/quickfix-back/pkg/repository/database/order"
	userRepo "github.com/jizambrana5/quickfix-back/pkg/repository/database/user"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
	"github.com/jizambrana5/quickfix-back/pkg/utils"
)

func main() {
	// Initialize the logger based on the environment
	logs.InitLogger("development")
	defer logs.Logger.Sync() //nolint

	// Ruta relativa al archivo JSON de ubicaciones en Mendoza
	locations, err := utils.LoadLocations()
	if err != nil {
		fmt.Println("error al cargar las locaciones", err.Error())
	}
	fmt.Println(locations)
	// Inicializar la conexión a la base de datos
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

	orderRepo := orderRepo.NewOrderRepository(db)
	orderSrv := order.NewService(orderRepo)

	userRepo := userRepo.NewUserRepository(db)
	userSrv := user.NewUserService(userRepo)

	handler := rest.NewHandler(orderSrv, userSrv)

	server := rest.Routes(handler)
	err = server.Run(fmt.Sprintf("%s%s", ":", "8080"))
	if err != nil {
		panic(err)
	}
}
