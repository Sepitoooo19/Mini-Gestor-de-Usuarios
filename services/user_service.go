package services

import (
	"context"
	"errors"
	"go-template/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UsersCollection es la referencia a la colección de usuarios en MongoDB
var UsersCollection *mongo.Collection

// InitUsersService inicializa la conexión a la colección de usuarios
func InitUsersService() {
	db := Client.Database("Mongo")
	UsersCollection = db.Collection("users")
}

// CreateUser crea un nuevo usuario en la base de datos
func CreateUser(user models.Users) error {
	if UsersCollection == nil {
		return errors.New("UsersCollection no está inicializada")
	}

	_, err := UsersCollection.InsertOne(context.Background(), user)
	return err
}

// GetUsers obtiene todos los usuarios de la base de datos
func GetUsers() ([]models.Users, error) {
	var users []models.Users

	if UsersCollection == nil {
		return nil, errors.New("UsersCollection no está inicializada")
	}

	cursor, err := UsersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID busca un usuario específico por su ID
func GetUserByID(id string) (models.Users, error) {
	var user models.Users

	if UsersCollection == nil {
		return user, errors.New("UsersCollection no está inicializada")
	}

	if id == "" {
		return user, errors.New("ID no puede estar vacío")
	}

	filter := bson.M{"_id": id}
	err := UsersCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("usuario no encontrado")
		}
		return user, err
	}

	return user, nil
}
