package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

//User Create a struct type to handle the user
type User struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Salt      string               `json:"salt,omitempty" bson:"salt,omitempty"`
	Email     string               `json:"email,omitempty" bson:"email,omitempty"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty"`
	Hash      []byte               `json:"hash,omitempty" bson:"hash,omitempty"`
	CreatedAt time.Time            `json:"createdat,omitempty" bson:"createdat,omitempty"`
	Threads   []primitive.ObjectID `json:"thread,omitempty" bson:"thread,omitempty"`
	IsAdmin   bool                 `json:"isadmin,omitempty" bson:"isadmin,omitempty"`
}

//LogInUser Create a struct type to handle the user
type LogInUser struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string               `json:"email,omitempty" bson:"email,omitempty"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt time.Time            `json:"createdat,omitempty" bson:"createdat,omitempty"`
	Threads   []primitive.ObjectID `json:"thread,omitempty" bson:"thread,omitempty"`
	IsAdmin   bool                 `json:"isadmin,omitempty" bson:"isadmin,omitempty"`
}
