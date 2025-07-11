package main

import (
	"github.com/gin-gonic/gin"
	"go-template/config"
	"go-template/routes"
	"go-template/services"
)

func main() {
	r := gin.Default()

	//Configurar el CORS para el llamado del front
	r.Use(config.SetupCORS())

	// Inicializar conexión Mongo
	services.InitMongo()

	// Inicializar servicio de usuarios
	services.InitUsersService()

	// Registrar rutas
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
