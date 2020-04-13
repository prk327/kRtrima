package models

import (
	"context"
	//	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindComment(res1 string, collection *mongo.Collection) *Comment {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// create a value into which the result can be decoded
	var result *Comment

	err = collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		Logger.Fatalln(err)
	}

	return result
}

func FindAllComment(limit int64, collection *mongo.Collection) []*Comment {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*Comment

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Comment
		err := cur.Decode(&elem)
		if err != nil {
			Logger.Fatalln(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		Logger.Fatalln(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}

func FindAllCommentByID(key string, res1 string, collection *mongo.Collection) []*Comment {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// Pass these options to the Find method
	findOptions := options.Find()
	//	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*Comment

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.M{key: docID}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Comment
		err := cur.Decode(&elem)
		if err != nil {
			Logger.Fatalln(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		Logger.Fatalln(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}
