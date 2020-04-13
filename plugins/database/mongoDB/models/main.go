package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var DB *mongo.Database

var Collection *mongo.Collection

var Comments *mongo.Collection

var Users *mongo.Collection

var Sessions *mongo.Collection

var Logger *log.Logger

var Msg string

var lm = &Msg

func init() {

	*lm, DB = Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
	fmt.Println(Msg)

	//    conect to Thread collection
	*lm, Collection = Cnt_Collection("Thread", DB)
	fmt.Println(Msg)

	//    conect to Comment collection
	*lm, Comments = Cnt_Collection("Comment", DB)
	fmt.Println(Msg)
    
    //    conect to User collection
	*lm, Users = Cnt_Collection("User", DB)
	fmt.Println(Msg)
    
    //    conect to Session collection
	*lm, Sessions = Cnt_Collection("Session", DB)
	fmt.Println(Msg)

}
