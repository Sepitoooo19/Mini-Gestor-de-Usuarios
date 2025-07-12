package main

import (
	"github.com/gin-gonic/gin"
	"go-template/config"
	"go-template/routes"
	"go-template/services"
)

func main() {
	r := gin.Default()

	// Configurar middleware CORS para permitir llamadas desde frontend
	r.Use(config.SetupCORS())

	// Inicializar conexión con MongoDB
	services.InitMongo()

	// Inicializar servicio de gestión de usuarios
	services.InitUsersService()

	// Registrar todas las rutas de la API
	routes.RegisterRoutes(r)

	// Iniciar servidor HTTP en puerto 8080
	r.Run(":8080")
}
