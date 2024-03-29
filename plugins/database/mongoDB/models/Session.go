package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

//Session Create a struct type to handle the session for login
type Session struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Salt      string             `json:"salt,omitempty" bson:"salt,omitempty"`
	CreatedAt time.Time          `json:"createdat" bson:"createdat"`
}
