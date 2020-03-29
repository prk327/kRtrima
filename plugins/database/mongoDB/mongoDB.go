package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// Create a struct type to handle the thread
type Thread struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Image       string
	Description string
}

var DB *mongo.Database

var Collection *mongo.Collection

var Logger *log.Logger

var Msg string