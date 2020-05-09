package comments

import (
	"github.com/julienschmidt/httprouter"

	m "kRtrima/plugins/database/mongoDB/models"

	"net/http"
)

// New is used to create a new comment
func New(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	err := m.Threads.Find("_id", p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to Find the Thread!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	// get the list of mongo collections
	coll, err := m.ShowCollectionNames(m.DB)
	if err != nil {
		Logger.Println("Not able to Get the list of Collection!!")
		http.Redirect(writer, request, "/", 302)
		return
	}

	dashlist := m.FindDetails{
		CollectionNames: coll,
		ContentDetails:  m.TP,
		User:            m.UP,
		LogInUser:       m.LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "CommentNew")
}
