package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home is to show the home page
func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

//Index is used to show the threads
func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	err := m.GetUserbyUUID("kRtrima", writer, request)
	if err != nil {
		Logger.Println("Not able to find the user")
		http.Redirect(writer, request, "/login", 401)
		return
	}

	//get the thread and assign it to slice of thread TSL
	err = m.Threads.FindAll(10)
	if err != nil {
		Logger.Println("Not able to Find the list of Thread!!")
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

	dashlist := m.MainCongifDetails{
		CollectionNames: coll,
		ContentDetails:  m.TSL,
		User:            m.UP,
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "index")
}
