package services

import (
	"context"
	"errors"
	"go-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// UsersCollection es la colección de usuarios en MongoDB
var UsersCollection *mongo.Collection

// InitUsersService inicializa la colección de usuarios en MongoDB
func InitUsersService() {
	db := Client.Database("Mongo")
	UsersCollection = db.Collection("users")
}

// CreateUser crea un nuevo usuario en la colección de usuarios
func CreateUser(user models.Users) error {
	// Verificar si la colección está inicializada
	if UsersCollection == nil {
		return errors.New("UsersCollection no está inicializada")
	}

	log.Printf("DEBUG - Creando usuario: ID=%s, Name=%s, Email=%s, CreatedAt=%v", 
		user.ID, user.Name, user.Email, user.CreatedAt)
	
	// Insertar el usuario en MongoDB
	result, err := UsersCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("ERROR - Falló la inserción en MongoDB: %v", err)
		return err
	}
	
	log.Printf("SUCCESS - Usuario insertado con ID: %v", result.InsertedID)
	return nil
}

// GetUsers obtiene todos los usuarios de la colección de usuarios
func GetUsers() ([]models.Users, error) {
	var users []models.Users

	// Verificar si la colección de usuarios está inicializada
	if UsersCollection == nil {
		return nil, errors.New("UsersCollection no está inicializada")
	}

	log.Printf("DEBUG - Obteniendo todos los usuarios")

	// Crear un cursor para iterar sobre los documentos en la colección de usuarios
	cursor, err := UsersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("ERROR - Error al obtener usuarios: %v", err)
		return nil, err
	}
	// Se cierra el cursor al final de la función para liberar recursos
	defer cursor.Close(context.Background())

	// Iterar sobre los documentos en el cursor y decodificarlos en la estructura de usuario
	for cursor.Next(context.Background()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			log.Printf("ERROR - Error al decodificar usuario: %v", err)
			continue // Continúa con el siguiente documento en lugar de fallar completamente
		}
		users = append(users, user)
	}

	// Verificar si hubo errores durante la iteración
	if err := cursor.Err(); err != nil {
		log.Printf("ERROR - Error durante la iteración del cursor: %v", err)
		return nil, err
	}

	log.Printf("SUCCESS - Se obtuvieron %d usuarios", len(users))
	return users, nil
}

// GetUserByID obtiene un usuario por su ID de la colección de usuarios
func GetUserByID(id string) (models.Users, error) {
	var user models.Users

	// Verificar si la colección de usuarios está inicializada
	if UsersCollection == nil {
		return user, errors.New("UsersCollection no está inicializada")
	}

	// Validar que el ID no esté vacío
	if id == "" {
		return user, errors.New("ID no puede estar vacío")
	}

	// Se realiza una búsqueda en la colección de usuarios utilizando el ID proporcionado
	filter := bson.M{"_id": id}

	log.Printf("DEBUG - Buscando usuario con ID: %s", id)

	// Se utiliza FindOne para buscar un único documento que coincida con el filtro
	err := UsersCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("DEBUG - Usuario no encontrado con ID: %s", id)
			return user, errors.New("usuario no encontrado")
		}
		log.Printf("ERROR - Error al buscar usuario: %v", err)
		return user, err
	}

	log.Printf("SUCCESS - Usuario encontrado: %s", user.Name)
	return user, nil
}
