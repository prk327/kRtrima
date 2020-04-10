package web

import (
	"github.com/julienschmidt/httprouter"
	"kRtrima/plugins/database/mongoDB"
	"net/http"
)

type DBConfig struct {
	Host       string
	Database   string
	Collection string
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

	_, mongoDB.Collection = mongoDB.Cnt_Collection(DB_Config.Collection, mongoDB.DB)

	http.Redirect(writer, request, "/", 302)

}