package controllers

import (
	"github.com/gin-gonic/gin"
	"go-template/models"
	"go-template/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateUser es el controlador para crear un nuevo usuario
func CreateUser(c *gin.Context) {
	var user models.Users

	// Bindear el JSON recibido en la petición al modelo de usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se asigna un ID único al usuario usando ObjectID de MongoDB
	user.ID = primitive.NewObjectID().Hex()

	err := services.CreateUser(user)
	// Si ocurre un error al crear el usuario, se retorna un mensaje de error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	// Si la creación es exitosa, se retorna el usuario creado con un código de estado 201
	c.JSON(http.StatusCreated, user)
}

// GetUsers es el controlador para obtener todos los usuarios
func GetUsers(c *gin.Context) {

	users, err := services.GetUsers()

	// Si ocurre un error al obtener los usuarios, se retorna un mensaje de error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}

	// Si la obtención es exitosa, se retorna la lista de usuarios con un código de estado 200
	c.JSON(http.StatusOK, users)
}

// GetUserByID es el controlador para obtener un usuario por su ID
func GetUserByID(c *gin.Context) {

	// Se obtiene el ID del usuario desde los parámetros de la ruta
	id := c.Param("id")

	// Convertir el ID a un ObjectID de MongoDB
	user, err := services.GetUserByID(id)

	// Si ocurre un error al obtener el usuario, se retorna un mensaje de error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Si la obtención es exitosa, se retorna el usuario con un código de estado 200
	c.JSON(http.StatusOK, user)
}
