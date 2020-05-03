package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Logger the variable to capture the log event
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("kRtrima_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "Model: ", log.Ldate|log.Ltime|log.Lshortfile)
}

//Msg the  variable to print the msg
var Msg string

// MongoConfig is used to store the MDB config
type MongoConfig struct {
	Hostname string          `json:"hostname,omitempty" bson:"hostname,omitempty"`
	DBName   string          `json:"dbname,omitempty" bson:"dbname,omitempty"`
	DB       *mongo.Database `json:"db,omitempty" bson:"db,omitempty"`
}

// MongoCollection is used to store the model configuration
type MongoCollection struct {
	CollectionName string
	Collection     *mongo.Collection
	Model          interface{}
	SLModel        interface{}
}

//NewMongoConnection initialization of mongodb
func NewMongoConnection(hostname string, dbaname string) *MongoConfig {

	mongoDBCo := MongoConfig{
		Hostname: hostname,
		DBName:   dbaname,
	}

	db, err := mongoDBCo.ConnectDB()
	if err != nil {
		Logger.Fatalln(err)
	}

	return &MongoConfig{
		Hostname: hostname,
		DBName:   dbaname,
		DB:       db,
	}
}

// ConnectDB is used to connect the Database
func (m MongoConfig) ConnectDB() (DB *mongo.Database, err error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(m.Hostname)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger.Println(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Println(err)
	}
	DB = client.Database(m.DBName)
	*&Msg = fmt.Sprintf("Connected to MongoDB Database %v at %v", m.DBName, m.Hostname)
	fmt.Println(Msg)
	return
}

// NewCollection is used to get the collection object and initialize the model
func (m *MongoConfig) NewCollection(Name string, model interface{}, SliceOFModel interface{}) (modelstruct *MongoCollection) {
	modelstruct = &MongoCollection{
		CollectionName: Name,
		Collection:     m.DB.Collection(Name),
		Model:          model,
		SLModel:        SliceOFModel,
	}
	*&Msg = fmt.Sprintf("Connected to MongoDB Collection: %v", Name)
	fmt.Println(Msg)
	return
}

// ShowCollectionNames is to display the list of all colection
func ShowCollectionNames(DB *mongo.Database) (result []string, err error) {
	// use a filter to only select capped collections
	result, err = DB.ListCollectionNames(context.TODO(), bson.D{{}})
	if err != nil {
		Logger.Println(err)
	}
	return
}

//AddItem To insert a single document working tested
func (m *MongoCollection) AddItem(t interface{}) (result primitive.ObjectID, err error) {
	insertResult, err := m.Collection.InsertOne(context.TODO(), t)
	if err != nil {
		Logger.Println(err)
	}
	result = primitive.ObjectID(insertResult.InsertedID.(primitive.ObjectID))
	return
}

// ToDocID this will return primitive object it change the type from string to objectID
func ToDocID(v interface{}) (DocID primitive.ObjectID, err error) {
	switch id := v.(type) {
	default:
		Logger.Fatalf("Unexpected Type %T\n", id)
	case primitive.ObjectID:
		DocID = primitive.ObjectID(v.(primitive.ObjectID))
	case string:
		re := regexp.MustCompile(`"(.*?)"`)
		rStr := fmt.Sprintf(`%v`, v)
		res1 := re.FindStringSubmatch(rStr)[1]
		// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
		docID, err := primitive.ObjectIDFromHex(res1)
		if err != nil {
			fmt.Printf("Cannot Convert %T type to object id", id)
			Logger.Println(err)
		}
		return docID, err
	}
	return
}

//UpdateItem To updated documents at a time
func (m *MongoCollection) UpdateItem(res1 interface{}, change interface{}) (retArg string, err error) {
	docID, err := ToDocID(res1)
	if err != nil {
		Logger.Println(err)
	}
	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"_id", docID}}
	update := bson.D{{"$set", change}}
	result, err := m.Collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		Logger.Fatal(err)
	}

	if result.MatchedCount != 0 {
		retArg = fmt.Sprintf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	}
	if result.UpsertedCount != 0 {
		retArg = fmt.Sprintf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	retArg = fmt.Sprintf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	return
}

// DeleteItem will delet any collection item
func (m *MongoCollection) DeleteItem(res1 interface{}) (result string, err error) {
	docID, err := ToDocID(res1)
	if err != nil {
		Logger.Println(err)
	}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	filter := bson.D{{"_id", docID}}
	res, err := m.Collection.DeleteOne(context.TODO(), filter, opts)
	if err != nil {
		Logger.Println(err)
	}
	result = fmt.Sprintf("deleted %v documents\n", res.DeletedCount)
	return
}

// DeleteAll will delet all the item in the given collection
func (m *MongoCollection) DeleteAll() (deletionMsg string, err error) {
	deleteResult, err := m.Collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		Logger.Println(err)
	}
	deletionMsg = fmt.Sprintf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return
}

// DropCollection will drop the collection
func (m *MongoCollection) DropCollection() (err error) {
	err = m.Collection.Drop(context.TODO())
	if err != nil {
		Logger.Println(err)
	}
	return
}

// Find will search for item using key value pair
func (m *MongoCollection) Find(key string, value interface{}) (err error) {
	var vx interface{}
	if key == "_id" {
		vx, err = ToDocID(value)
	} else {
		vx = value
	}
	if err != nil {
		Logger.Println(err)
	}
	err = m.Collection.FindOne(context.TODO(), bson.M{key: vx}).Decode(m.Model)
	if err != nil {
		Logger.Println(err)
	}
	return
}

// FindAll will return all the documents in a collection
func (m *MongoCollection) FindAll(limit int64) (err error) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cursor, err := m.Collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		Logger.Println(err)
	}
	if err = cursor.All(context.TODO(), *&m.SLModel); err != nil {
		Logger.Println(err)
	}
	cursor.Close(context.TODO())
	return
}

// FindbyKeyValue will return all the documents in a collection
func (m *MongoCollection) FindbyKeyValue(key string, value interface{}) (err error) {
	cursor, err := m.Collection.Find(context.TODO(), bson.M{key: value})
	if err != nil {
		Logger.Println(err)
	}
	if err = cursor.All(context.TODO(), *&m.SLModel); err != nil {
		Logger.Println(err)
	}
	cursor.Close(context.TODO())
	return
}
