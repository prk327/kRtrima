package models

import (
	"context"
	//	"flag"
	"fmt"
	"log"
	"time"

	// Official 'mongo-go-driver' packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IndexModel struct {
	cmd        string `json:"cmd,omitempty" bson:"cmd,omitempty"`
	address    string `json:"address,omitempty" bson:"address,omitempty"`
	database   string `json:"database,omitempty" bson:"database,omitempty"`
	collection string `json:"collection,omitempty" bson:"collection,omitempty"`
	key        string `json:"key,omitempty" bson:"key,omitempty"`
	unique     bool   `json:"unique,omitempty" bson:"unique,omitempty"`
	value      int    `json:"value,omitempty" bson:"value,omitempty"`
}

//"cmd", "", "list or add?"
//"address", "", "mongodb address to connect to"
//"db", "", "The name of the database to connect to"
//"collection", "", "The collection (in the db) to connect to"
//"key", "", "The field you'd like to place an index on"
//"unique", false, "Would you like the index to be unique?"
//"value", 1, "would you like the index to be ascending (1) or descending (-1)?"
func (i IndexModel) CreateIndex() {
	switch {
	case (i.cmd != "add" && i.cmd != "list"):
		log.Fatalf("The first argument has to be 'add' or 'list :)")
	case i.database == "" || i.address == "":
		log.Fatalf("Please provide a valid database address and database name :)")

	case i.cmd == "add" && i.key == "":
		log.Fatalf("Please pass in the name of the field to place the index :)")
	}
	client := i.ConnectToTheMongoDB()
	if i.cmd == "add" {
		i.PopulateIndex(client)
	} else if i.cmd == "list" {
		i.ListIndexes(client)
	}
}

func (i IndexModel) ConnectToTheMongoDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(i.address)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully connected to the address provided")
	return client
}

func (i IndexModel) PopulateIndex(client *mongo.Client) {
	c := client.Database(i.database).Collection(i.collection)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	index := i.yieldIndexModel()
	c.Indexes().CreateOne(context.Background(), index, opts)
	log.Println("Successfully create the index")
}

func (i IndexModel) yieldIndexModel() mongo.IndexModel {
	keys := bsonx.Doc{
		{
			Key:   i.key,
			Value: bsonx.Int32(int32(i.value)),
		},
	}
	index := mongo.IndexModel{}
	index.Keys = keys
	if i.unique {
		index.Options = options.Index().SetUnique(true)
	}
	return index
}

func (i IndexModel) ListIndexes(client *mongo.Client) {
	c := client.Database(i.database).Collection(i.collection)
	duration := 10 * time.Second
	batchSize := int32(10)
	cur, err := c.Indexes().List(context.Background(), &options.ListIndexesOptions{&batchSize, &duration})
	if err != nil {
		log.Fatalf("Something went wrong listing %v", err)
	}
	for cur.Next(context.Background()) {
		index := bson.D{}
		cur.Decode(&index)
		log.Println(fmt.Sprintf("index found %v", index))
	}
}
