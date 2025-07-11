package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/models"
	"go-template/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// CreateUser es el controlador para crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.Users

	// Bindear el JSON recibido en la petición al modelo de usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos: " + err.Error()})
		return
	}

	// Validar campos requeridos
	if user.Name == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y email son campos obligatorios"})
		return
	}

	// Asignar un ID único al usuario usando ObjectID de MongoDB
	user.ID = primitive.NewObjectID().Hex()

	// Asignar la fecha de creación actual del sistema (SIEMPRE sobreescribir)
	user.CreatedAt = time.Now().UTC()

	// Log de depuración
	fmt.Printf("DEBUG - Usuario a crear: ID=%s, Name=%s, Email=%s, CreatedAt=%v\n", 
		user.ID, user.Name, user.Email, user.CreatedAt)

	// Llamar al servicio para crear el usuario
	err := services.CreateUser(user)
	if err != nil {
		fmt.Printf("ERROR - No se pudo crear el usuario: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	// Log de éxito
	fmt.Printf("SUCCESS - Usuario creado correctamente: %s\n", user.ID)

	// Respuesta con el usuario creado
	c.JSON(http.StatusCreated, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"birthdate":  user.Birthdate,
		"created_at": user.CreatedAt.Format(time.RFC3339),
	})
}

// GetUsers es el controlador para obtener todos los usuarios
func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()

	// Si ocurre un error al obtener los usuarios, se retorna un mensaje de error
	if err != nil {
		fmt.Printf("ERROR - No se pudieron obtener los usuarios: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}

	// Formatear la respuesta - incluir todos los usuarios (incluso con created_at vacío)
	var formattedUsers []gin.H
	for _, user := range users {
		formattedUser := gin.H{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"birthdate": user.Birthdate,
		}
		
		// Solo agregar created_at si no está vacío
		if !user.CreatedAt.IsZero() {
			formattedUser["created_at"] = user.CreatedAt.Format(time.RFC3339)
		} else {
			formattedUser["created_at"] = nil
		}
		
		formattedUsers = append(formattedUsers, formattedUser)
	}

	fmt.Printf("SUCCESS - Se retornaron %d usuarios\n", len(formattedUsers))

	// Retornar solo la lista de usuarios (sin total)
	c.JSON(http.StatusOK, formattedUsers)
}

// GetUserByID es el controlador para obtener un usuario por su ID
func GetUserByID(c *gin.Context) {
	// Se obtiene el ID del usuario desde los parámetros de la ruta
	id := c.Param("id")

	// Validar que el ID no esté vacío
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	fmt.Printf("DEBUG - Buscando usuario con ID: %s\n", id)

	// Obtener el usuario por ID desde el servicio
	user, err := services.GetUserByID(id)

	// Si ocurre un error al obtener el usuario, se retorna un mensaje de error
	if err != nil {
		fmt.Printf("ERROR - Usuario no encontrado: %v\n", err)
		if err.Error() == "usuario no encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	fmt.Printf("SUCCESS - Usuario encontrado: %s\n", user.Name)

	// Formatear la respuesta para ser consistente con CreateUser
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"birthdate":  user.Birthdate,
		"created_at": user.CreatedAt.Format(time.RFC3339),
	})
}
