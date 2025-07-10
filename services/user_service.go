package services

import (
	"context"
	"errors"
	"go-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	_, err := UsersCollection.InsertOne(context.Background(), user)
	return err
}

// GetUsers obtiene todos los usuarios de la colección de usuarios
func GetUsers() ([]models.Users, error) {
	var users []models.Users

	// Verificar si la colección de usuarios está inicializada
	if UsersCollection == nil {
		return nil, errors.New("UsersCollection no está inicializada")
	}

	// Crear un cursor para iterar sobre los documentos en la colección de usuarios
	cursor, err := UsersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	// Se cierra el cursor al final de la función para liberar recursos
	defer cursor.Close(context.Background())

	// Iterar sobre los documentos en el cursor y decodificarlos en la estructura de usuario
	for cursor.Next(context.Background()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	//Retornar la lista de usuarios si no hay errores, si hay errores, se retornará nil
	return users, nil
}

// GetUserByID obtiene un usuario por su ID de la colección de usuarios
func GetUserByID(id string) (models.Users, error) {
	var user models.Users

	// Verificar si la colección de usuarios está inicializada
	if UsersCollection == nil {
		return user, errors.New("UsersCollection no está inicializada")
	}

	// Se realiza una búsqueda en la colección de usuarios utilizando el ID proporcionado
	filter := bson.M{"_id": id}

	// Se utiliza FindOne para buscar un único documento que coincida con el filtro
	err := UsersCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return user, err
	}

	// Si se encuentra el usuario, se retorna; de lo contrario, se retorna un error
	return user, nil
}
