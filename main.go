package main

import (
	"kRtrima/plugins/database/mongoDB"
	"kRtrima/web"
    "fmt"
)

func main() {
    //running the mongoDB server
    msg, _ := mongoDB.Run_mongoDB("mongodb://localhost:27017", "kRtrima", "Thread")
    
    fmt.Println(msg)
	//running the kRtrima web server
	web.Web()
}
