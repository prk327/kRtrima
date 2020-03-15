package main

import (
	"kRtrima/plugins/database/mongoDB"
	"kRtrima/web"
)

func main() {
    //running the mongoDB server
	mongoDB.Run_mongoDB()
	//running the kRtrima web server
	web.Web()
}
