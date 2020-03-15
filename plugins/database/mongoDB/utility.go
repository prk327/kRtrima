package mongoDB

import (
	"context"
	"fmt"
	//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Create a struct type to handle the thread
type Thread struct {
	Name  string
	Image string
}

//dummy data for testing
var p1 = Thread{"Circuit", "/resources/images/circuit.jpg"}
var p2 = Thread{"City At Night", "/resources/images/CityNight.jpg"}
var p3 = Thread{"Pyramid", "/resources/images/piramid.jpg"}

//To insert a single document working tested
func AddItem(t Thread, collection *mongo.Collection) string {
	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Inserted a single document: %v", insertResult.InsertedID)
}

//To insert multiple documents at a time
//func AddItems(t []interface {}, collection *mongo.Collection) string {
//	insertManyResult, err := collection.InsertMany(context.TODO(), t)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return fmt.Sprintf("Inserted multiple documents: %v", insertManyResult.InsertedIDs)
//}
//
//func UpdateItem(filter interface{}, method string, item primitive.E, collection *mongo.Collection) string {
//	//filter option
//	f := bson.D{filter}
//	//value to updated
//	update := bson.D{{method, bson.D{item}}}
//	updateResult, err := collection.UpdateOne(context.TODO(), f, update)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return fmt.Sprintf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
//}
//
//func FindItem(filter primitive.E, collection *mongo.Collection) Thread {
//	// create a value into which the result can be decoded
//	var result Thread
//
//	err = collection.FindOne(context.TODO(), filter).Decode(&result)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return result
//}
//
//func FindAllItem(limit int64, collection *mongo.Collection) string {
//	// Pass these options to the Find method
//	findOptions := options.Find()
//	findOptions.SetLimit(limit)
//
//	// Here's an array in which you can store the decoded documents
//	var results []*Thread
//
//	// Passing bson.D{{}} as the filter matches all documents in the collection
//	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Finding multiple documents returns a cursor
//	// Iterating through the cursor allows us to decode documents one at a time
//	for cur.Next(context.TODO()) {
//
//		// create a value into which the single document can be decoded
//		var elem Thread
//		err := cur.Decode(&elem)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		results = append(results, &elem)
//	}
//
//	if err := cur.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Close the cursor once finished
//	cur.Close(context.TODO())
//
//    return fmt.Sprintf("Found multiple documents (array of pointers): %+v\n", results)
//}
//
