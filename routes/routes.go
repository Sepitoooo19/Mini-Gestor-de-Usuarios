package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/controllers"
)

// RegisterRoutes configura todas las rutas HTTP disponibles en la aplicación.
func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.PingController)
	
	// Rutas de gestión de usuarios,
	router.POST("/users", controllers.CreateUser)      // Crear usuario
	router.GET("/users", controllers.GetUsers)         // Obtener todos los usuarios
	router.GET("/users/:id", controllers.GetUserByID)  // Obtener usuario por ID
}
