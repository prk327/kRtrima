package mongoDB

import (
	"context"
	"fmt"
	"log"
	//    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run_mongoDB(hostname string, databaseName string, documents string) (string, *mongo.Collection) {
	// Set client options
	clientOptions := options.Client().ApplyURI(hostname)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(databaseName).Collection(documents)

	//some useful utility functions
	return fmt.Sprintf("Connected to MongoDB database %v at %v for collection %v", databaseName, hostname, documents), collection

}
