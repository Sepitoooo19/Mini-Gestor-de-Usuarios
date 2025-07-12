package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// SetupCORS configura el middleware CORS para permitir solicitudes desde el frontend
func SetupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true, // En producción, cambiar por dominios específicos
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
