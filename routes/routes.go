package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.PingController)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
}
