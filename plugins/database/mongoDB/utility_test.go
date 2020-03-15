package mongoDB

//import (
//	"context"
//	"log"
//	"testing"
//	//    "go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func TestAddItem(t *testing.T) {
//	// Set client options
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//
//	// Connect to MongoDB
//	client, err := mongo.Connect(context.TODO(), clientOptions)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Check the connection
//	err = client.Ping(context.TODO(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	collection := client.Database("kRtrima").Collection("Thread")
//
//	want := "Inserted a single document"
//	if got := AddItem(p2, collection); got != want {
//		t.Errorf("AddItem() = %q, want %q", got, want)
//	}
//}
