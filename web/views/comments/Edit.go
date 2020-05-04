package comments

import (
	"github.com/julienschmidt/httprouter"

	m "kRtrima/plugins/database/mongoDB/models"

	"net/http"
)

// Edit is used to edit the thread
func Edit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	//find the thread by id and assign it to TP
	err := m.Threads.Find("_id", p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to Find the thread by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	//find the comment by id and assign it to CSL
	err = m.Comments.Find("_id", p.ByName("cid"))
	if err != nil {
		Logger.Println("Not able to Find the comment by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	//get the User and assign to User UP struct
	err = m.Users.Find("_id", m.CP.User)
	if err != nil {
		Logger.Println("Not able to Find the user by ID!!")
		http.Redirect(writer, request, "/", 302)
		return
	}

	// get the list of mongo collections
	coll, err := m.ShowCollectionNames(m.DB)
	if err != nil {
		Logger.Println("Not able to Get the list of Collection!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	dashlist := m.FindDetails{
		CollectionNames: coll,
		ContentDetails:  m.TP,
		Comments:        m.CSL,
		User:            m.UP,
		LogInUser:       m.LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "CommentEdit")
}
