package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/repository/database"
	"github.com/jizambrana5/quickfix-back/pkg/rest"
)

func main() {
	r := gin.Default()

	// Inicializar la conexión a la base de datos
	database.InitDatabase()

	// Registrar rutas
	rest.Routes(r)

	// Iniciar el servidor
	err := r.Run(fmt.Sprintf("%s%s", ":", "8080"))
	if err != nil {
		panic(err)
	}
}
