package models

import (
	"context"
	//	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    "time"
//    "strings"
	//	"log"
)

func FindUserByID(res1 string, collection *mongo.Collection) *User {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}

	var result *User

	err = collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		Logger.Fatalln(err)
	}

	return result
}

func FindUser(key string, value string, collection *mongo.Collection) *User {

	var result *User

	err := collection.FindOne(context.TODO(), bson.M{key: value}).Decode(&result)
	if err != nil {
		Logger.Fatalln(err)
	}

	return result
}

func FindAllUser(limit int64, collection *mongo.Collection) []*User {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
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


func FindAllUserByID(key string, res1 string, collection *mongo.Collection) []*User {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// Pass these options to the Find method
	findOptions := options.Find()
	//	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.M{key: docID}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
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


// Create a new session for an existing user
func (user *User) CreateSession() (stmt *Session, err error) {
    
     // Create a struct type to handle the session for login
    statement :=  Session {
        Uuid: CreateUUID(),
        Email: user.Email,
        UserId: user.ID,
        CreatedAt: time.Now(),
}
 
    ssid, err := AddItem(statement, Sessions)
    if err != nil {
			Logger.Fatalln(err)
		}
    
    stmt, err = FindSessionByID(ssid, Sessions)
    
    if err != nil {
			Logger.Fatalln(err)
		}
    
	
	return
}
