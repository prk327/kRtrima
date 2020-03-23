package mongoDB

import (
	//	"context"
	//	"fmt"
	//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	//	"os"
)

// Create a struct type to handle the thread
type Thread struct {
	Name  string
	Image string
}

var DB *mongo.Database

var Collection *mongo.Collection

//var postData Thread

var Logger *log.Logger

//dummy data for testing
var P1 = Thread{"Circuit", "/resources/images/circuit.jpg"}

//var p2 = Thread{"City At Night", "/resources/images/CityNight.jpg"}
//var p3 = Thread{"Pyramid", "/resources/images/piramid.jpg"}
var Msg string

//func MongoDBMain() {
//	lm := &msg
//	//    connect to db
//	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
//	logger.Println(msg)
//	//    conect to collection
//	*lm, Collection = Cnt_Collection("Thread", DB)
//	logger.Println(msg)
//
//	//    add item
//	//	*lm = AddItem(P1, Collection)
//	//	logger.Println(msg)
//
//	//    list the name of the collection
//	//	res := ShowCollectionNames(DB)
//	//	for _, coll := range res {
//	//		logger.Println(coll)
//	//	}
//
//	//    find all items in a thread
//	items := FindAllItem(10, Collection)
//	for _, coll := range items {
//		logger.Println(coll)
//	}
//}
