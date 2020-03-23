package web

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kRtrima/plugins/database/mongoDB"
	"net/http"
)


var lm = &mongoDB.Msg

//var Mcoll *mongo.Collection
//running the mongoDB server
//var DB *mongo.Database
//var hostName string
//var DBName string
//var collName string

type mainCongifDetails struct {
	collectionNames []string
	contentDetails  []*mongoDB.Thread
}

type DBConfig struct {
	Host       string
	Database   string
	Collection string
}

//     fmt.Println(msg)

//func init(){
//    msg, Mcoll := mongoDB.Run_mongoDB("mongodb://localhost:27017", "kRtrima").Collection("Thread")
//     fmt.Println(msg)
//}

//    DB.Collection("Thread")

//dummy data for testing
//var p1 = map[string]string{
//	"Name":  "Circuit",
//	"Image": "/resources/images/circuit.jpg",
//}
//var p2 = map[string]string{
//	"Name":  "City At Night",
//	"Image": "/resources/images/CityNight.jpg",
//}
//var p3 = map[string]string{
//	"Name":  "Pyramid",
//	"Image": "/resources/images/piramid.jpg",
//}
//var dashList = []map[string]string{p1, p2, p3}

func dashboard(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
    
    *lm, mongoDB.DB = mongoDB.Connect_mongoDB("mongodb://localhost:27017", "kRtrima")
    mongoDB.Logger.Println(mongoDB.Msg)

	//    conect to collection
	*lm, mongoDB.Collection = mongoDB.Cnt_Collection("Thread", mongoDB.DB)
	mongoDB.Logger.Println(mongoDB.Msg)

//	thread := mongoDB.FindAllItem(10, mongoDB.Collection)
//    
//    fmt.Printf("The type of thread is: %T\n", thread)
//    
//
//    result := mongoDB.ShowCollectionNames(mongoDB.DB)
//    
//    fmt.Printf("The type of result is is: %T\n", result)

    dashlist := mainCongifDetails{
        collectionNames: mongoDB.ShowCollectionNames(mongoDB.DB),
        contentDetails: mongoDB.FindAllItem(10, mongoDB.Collection),
	}
    
    	for _, coll := range dashlist.collectionNames {
		fmt.Println(coll)
	}
	
	generateHTML(writer, &dashlist.contentDetails, "layout", "leftsidebar", "topsidebar", "modal", "dashboard")
}

func postData(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	//  name := request.Form["name"]
	//	image := request.Form["image"]

	//	newItem := mongoDB.Thread{
	//		request.Form["name"][0],
	//		request.Form["image"][0],
	//	}

	//	mongoDB.AddItem(newItem, mongoDB.Collection)
	//	name := request.Form["name"]
	//	image := request.Form["image"]
	//	p4 := map[string]string{
	//		"Name":  name[0],
	//		"Image": image[0],
	//	}

	//	dashList = append(dashList, p4)

	http.Redirect(writer, request, "/", 302)

	//	generateHTML(writer, dashList, "layout", "leftsidebar", "topsidebar", "modal", "dashboard")
}

func landingPage(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func dataForm(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}

func connectDB(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	DB_Config := DBConfig{
		request.Form["mongo_HostName"][0],
		request.Form["mongo_DBName"][0],
		request.Form["mongo_CollName"][0],
	}

	_, mongoDB.DB = mongoDB.Connect_mongoDB(DB_Config.Host, DB_Config.Database)

//	fmt.Println(msg)

	_, mongoDB.Collection = mongoDB.Cnt_Collection(DB_Config.Collection, mongoDB.DB)

//	fmt.Println(msgA)

	http.Redirect(writer, request, "/", 302)

}
