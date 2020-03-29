package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("plugins/database/mongoDB/mongoDB_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

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

func ShowCollectionNames(DB *mongo.Database) []string {
	// use a filter to only select capped collections
	result, err := DB.ListCollectionNames(context.TODO(), bson.D{{}})
	if err != nil {
		Logger.Fatal(err)
		//        return
	}

	return result
}

func Cnt_Collection(Name string, DB *mongo.Database) (string, *mongo.Collection) {
	collection := DB.Collection(Name)
	return fmt.Sprintf("Connected to MongoDB Collection: %v", Name), collection
}

//func UpdateCollection(method string, key string, value interface{}) bson.D {
//	update := bson.D{
//		{method, bson.D{
//			{key, value},
//		}},
//	}
//	return update
//}

//To insert a single document working tested
func AddItem(t Thread, collection *mongo.Collection) string {
	insertResult, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		Logger.Fatalln(err)
	}
	return fmt.Sprintf("Inserted a single document: %v", insertResult.InsertedID)
}

////To insert multiple documents at a time
//func AddItems(t []interface{}, collection *mongo.Collection) string {
//	insertManyResult, err := collection.InsertMany(context.TODO(), t)
//	if err != nil {
//		//        return fmt.Sprintf("Got an Error %v while saving the doc", err)
//		logger.Fatalln(err)
//
//	}
//	return fmt.Sprintf("Inserted multiple documents: %v", insertManyResult.InsertedIDs)
//}

//To updated documents at a time
//func UpdateItem(filter bson.D, method string, update bson.D, collection *mongo.Collection) string {
//	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
//	if err != nil {
//		logger.Fatalln(err)
//	}
//	return fmt.Sprintf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
//}

func FindItem(docID primitive.ObjectID, collection *mongo.Collection) *Thread {
	// create a value into which the result can be decoded
	var result *Thread

	err := collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		Logger.Fatalln(err)
	}

	return result
}

func FindAllItem(limit int64, collection *mongo.Collection) []*Thread {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// Here's an array in which you can store the decoded documents
	var results []*Thread

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		Logger.Fatalln(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Thread
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
