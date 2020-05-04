package models

import (
	// "fmt"
	"go.mongodb.org/mongo-driver/mongo"
	// "log"
)

// DB is a variable to store the connection
var DB *mongo.Database

// Threads is a variable to store MongoCollection
var Threads *MongoCollection

//TP assign the pointer to the Thread model
var TP *Thread

//TSL assign the Slice of Thread
var TSL []Thread

// Comments is a variable to store MongoCollection
var Comments *MongoCollection

//CP assign the pointer to the Comment model
var CP *Comment

//CSL assign the Slice of Comment
var CSL []Comment

// Users is a variable to store MongoCollection
var Users *MongoCollection

//UP assign the pointer to the User model
var UP *User

//USL assign the Slice of USER
var USL []User

// LogInUser is a variable to store MongoCollection
var LogInUser *MongoCollection

//LIP assign the pointer to the User model
var LIP *User

//LISL assign the Slice of USER
var LISL []User

//-------------

// Sessions is a variable to store MongoCollection
var Sessions *MongoCollection

//SP assign the pointer to the Session model
var SP *Session

//SSL assign the Slice of Session
var SSL []Session

func init() {

	db := NewMongoConnection("mongodb://localhost:27017", "kRtrima")

	DB = db.DB

	//Get the pointer to mongo Collection for Thread
	Threads = db.NewCollection("Thread", &TP, &TSL)

	//conect to Comment collection
	Comments = db.NewCollection("Comment", &CP, &CSL)

	//conect to User collection
	Users = db.NewCollection("User", &UP, &USL)

	//conect to User collection
	LogInUser = db.NewCollection("User", &LIP, &LISL)

	//conect to Session collection
	Sessions = db.NewCollection("Session", &SP, &SSL)

}
