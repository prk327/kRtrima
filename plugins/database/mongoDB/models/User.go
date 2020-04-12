package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

// Create a struct type to handle the user
type User struct {
	ID      primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Email   string               `json:"email,omitempty" bson:"email,omitempty"`
	Name    string               `json:"name,omitempty" bson:"name,omitempty"`
	Threads []primitive.ObjectID `json:"thread,omitempty" bson:"thread,omitempty"`
}
