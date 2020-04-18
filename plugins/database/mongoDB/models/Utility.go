package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
//    "strings"
)

func Connect_mongoDB(hostname string, databaseName string) (string, *mongo.Database) {
	// Set client options
	clientOptions := options.Client().ApplyURI(hostname)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger.Fatalln(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Fatalln(err)
	}
	return fmt.Sprintf("Connected to MongoDB database %v at %v", databaseName, hostname), client.Database(databaseName)
}

func Cnt_Collection(Name string, DB *mongo.Database) (string, *mongo.Collection) {
	collection := DB.Collection(Name)
	return fmt.Sprintf("Connected to MongoDB Collection: %v", Name), collection
}

func ShowCollectionNames(DB *mongo.Database) []string {
	// use a filter to only select capped collections
	result, err := DB.ListCollectionNames(context.TODO(), bson.D{{}})
	if err != nil {
		Logger.Fatal(err)
		//        return
	}

	return result
}

//To insert a single document working tested
func AddItem(t interface{}, collection *mongo.Collection) (result primitive.ObjectID, err error) {
    insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		Logger.Fatalln(err)
	}
    
    result = primitive.ObjectID(insertResult.InsertedID.(primitive.ObjectID))
    
	return
}

//To updated documents at a time
func UpdateItem(res1 string, change interface{}, collection *mongo.Collection) (retArg string, err error) {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
    docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// find the document for which the _id field matches id
	// specify the Upsert option to insert a new document if a document matching the filter isn't found
	//    uFil := primitive.E(change)
	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"_id", docID}}
	update := bson.D{{"$set", change}}
	//    fmt.Println(bson.D{uFil})
    result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		Logger.Fatal(err)
	}
    

	if result.MatchedCount != 0 {
        retArg = fmt.Sprintf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	}
	if result.UpsertedCount != 0 {
		retArg  = fmt.Sprintf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	retArg = fmt.Sprintf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
    return
}

func DeleteItem(res1 string, collection *mongo.Collection) (result string, err error) {
    
    // Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		Logger.Fatalln(err)
	}
    
	// delete at most one document in which the "name" field is "Bob" or "bob"
	// specify the SetCollation option to provide a collation that will ignore case for string comparisons
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	filter := bson.D{{"_id", docID}}
    res, err := collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		Logger.Fatal(err)
	}
    result = fmt.Sprintf("deleted %v documents\n", res.DeletedCount)
	return 
}

func DeleteAll(collection *mongo.Collection) string {
    deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
    if err != nil {
        Logger.Fatal(err)
    }
    return fmt.Sprintf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func DropCollection(collection *mongo.Collection) string {
    err := collection.Drop(context.TODO())
    if err != nil {
        Logger.Fatal(err)
    }
    return fmt.Sprintf("Deleted the collection!!")
}


func Findmodel(key string, value interface{}, collection *mongo.Collection) (result []bson.M, err error) {
    filterCursor, err := collection.Find(context.TODO(), bson.M{key: value})
    if err != nil {
        Logger.Fatal(err)
    }
    if err = filterCursor.All(context.TODO(), &result); err != nil {
        Logger.Fatal(err)
    }
	return
}
