package main

import (
//	"fmt"
//	"kRtrima/plugins/database/mongoDB"
	"kRtrima/web"
)

func main() {

	//    msg, DB := mongoDB.Run_mongoDB("mongodb://localhost:27017", "kRtrima")
	//
	//    DB.Collection("Thread")

	//running the kRtrima web server
	web.Web()
}
