package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"time"
)

// Create a struct type to handle the user
type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Uuid      string               `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Email     string               `json:"email" bson:"email"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty"`
	Password  []byte               `json:"password" bson:"password"`
	CreatedAt time.Time            `json:"createdat" bson:"createdat"`
	Threads   []primitive.ObjectID `json:"thread,omitempty" bson:"thread,omitempty"`
}
