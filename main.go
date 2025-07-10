package main

import (
	"github.com/gin-gonic/gin"
	"go-template/routes"
	"go-template/services"
)

func main() {
	// Inicializar conexión Mongo
	services.InitMongo()

	// Inicializar servicio de usuarios
	services.InitUsersService()

	r := gin.Default()

	// Registrar rutas
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
