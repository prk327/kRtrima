package models

import (
	//	"context"
	"fmt"
	//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
//	"os"
)

var DB *mongo.Database

var Collection *mongo.Collection

var Comments *mongo.Collection

var Logger *log.Logger

var Msg string

var lm = &Msg

func init() {

//	file, err := os.OpenFile("../models_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalln("Failed to open log file", err)
//	}
//	Logger = log.New(file, "Model: ", log.Ldate|log.Ltime|log.Lshortfile)

	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	fmt.Println(Msg)

	//    conect to Thread collection
	*lm, Collection = Cnt_Collection("Thread", DB)
	fmt.Println(Msg)

	//    conect to Comment collection
	*lm, Comments = Cnt_Collection("Comment", DB)
	fmt.Println(Msg)
}
