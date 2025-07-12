package controllers

import (
	"github.com/gin-gonic/gin"
	"go-template/models"
	"go-template/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// CreateUser maneja las peticiones POST para crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos: " + err.Error()})
		return
	}

	if user.Name == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y email son campos obligatorios"})
		return
	}

	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now().UTC()

	err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"birthdate":  user.Birthdate,
		"created_at": user.CreatedAt.Format(time.RFC3339),
	})
}

// GetUsers maneja las peticiones GET para obtener todos los usuarios
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}

	var formattedUsers []gin.H
	for _, user := range users {
		formattedUser := gin.H{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"birthdate": user.Birthdate,
		}
		
		if !user.CreatedAt.IsZero() {
			formattedUser["created_at"] = user.CreatedAt.Format(time.RFC3339)
		} else {
			formattedUser["created_at"] = nil
		}
		
		formattedUsers = append(formattedUsers, formattedUser)
	}

	c.JSON(http.StatusOK, formattedUsers)
}

// GetUserByID maneja las peticiones GET para obtener un usuario específico por ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	user, err := services.GetUserByID(id)

	if err != nil {
		if err.Error() == "usuario no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"birthdate":  user.Birthdate,
		"created_at": user.CreatedAt.Format(time.RFC3339),
	})
}
