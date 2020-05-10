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

	// //get the thread and assign to Thread TP struct
	err := m.Threads.Find("_id", p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to Find the thread by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	//delete the thread by id
	_, err = m.Threads.DeleteItem(p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to add new thread to DB!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	//get the comment and assign to Comment CP struct
	err = m.Comments.FindbyKeyValue("thread", m.TP.ID)
	if err != nil {
		Logger.Println("Not able to Find The Comments by thread ID!!")
	}

	// delete all the comments of the thread
	for _, c := range m.CSL {
		_, err = m.Comments.DeleteItem(c.ID)
		if err != nil {
			Logger.Println("Not able to delete comments for this thread!!")
			http.Redirect(writer, request, "/Dashboard", 302)
			return
		}
	}

	http.Redirect(writer, request, "/Dashboard", 302)
}
