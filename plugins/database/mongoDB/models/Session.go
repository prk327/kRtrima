package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"time"
)

// Create a struct type to handle the session for login
type Session struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Uuid      string             `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Email     string             `json:"email" bson:"email"`
	UserId    primitive.ObjectID `json:"userid" bson:"userid"`
	CreatedAt time.Time          `json:"createdat" bson:"createdat"`
}
