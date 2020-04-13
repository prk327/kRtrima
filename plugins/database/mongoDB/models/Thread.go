package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

// Create a struct type to handle the thread
type Thread struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name,omitempty" bson:"name,omitempty"`
	Image       string               `json:"image,omitempty" bson:"image,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	Comments    []primitive.ObjectID `json:"comments,omitempty" bson:"comments,omitempty"`
	User        primitive.ObjectID   `json:"user,omitempty" bson:"user,omitempty"`
}

type ThreadComment struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Comments    Comment            `json:"comments,omitempty" bson:"comments,omitempty"`
	User        primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
}
