package models

import "time"

type Users struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Birthdate string    `json:"birthdate" bson:"birthdate"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
