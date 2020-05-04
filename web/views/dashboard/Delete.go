package dashboard

import (
	// "fmt"
	m "kRtrima/plugins/database/mongoDB/models"

	"github.com/julienschmidt/httprouter"

	//    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"net/http"
	//	"regexp"
)

// Delete is used to delete a particular item
func Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	//delete the thread by id
	_, err := m.Threads.DeleteItem(p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to add new thread to DB!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	http.Redirect(writer, request, "/Dashboard", 302)
}
