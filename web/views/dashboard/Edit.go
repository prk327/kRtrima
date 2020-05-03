package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Edit is used to edit the thread
func Edit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	//find the thread by id and assign it to TP
	err := m.Threads.Find("_id", p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to Find the thread by ID!!")
		http.Redirect(writer, request, "/home", 401)
		return
	}

	// get the list of mongo collections
	coll, err := m.ShowCollectionNames(m.DB)
	if err != nil {
		Logger.Println("Not able to Get the list of Collection!!")
		http.Redirect(writer, request, "/", 301)
		return
	}

	dashlist := m.FindDetails{
		CollectionNames: coll,
		ContentDetails:  m.TP,
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "editData")
}
