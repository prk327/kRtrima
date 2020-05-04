package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

//Comment Create a struct type to handle the comments
type Comment struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Comment   string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Thread    primitive.ObjectID `json:"thread,omitempty" bson:"thread,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	CreatedAt time.Time          `json:"createdat" bson:"createdat"`
}
