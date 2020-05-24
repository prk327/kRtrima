package comments

import (
	"github.com/julienschmidt/httprouter"

	m "kRtrima/plugins/database/mongoDB/models"

	"net/http"
)

// New is used to create a new comment
func New(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	var TP m.Thread

	err := m.Threads.Find("_id", p.ByName("id"), &TP)
	if err != nil {
		Logger.Println("Not able to Find the Thread!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	var UP m.User

	err = m.Users.Find("_id", TP.User, &UP)
	if err != nil {
		Logger.Println("Not able to Find the User!!")
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

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}

	dashlist := m.FindDetails{
		CollectionNames: coll,
		ContentDetails:  &TP,
		User:            &UP,
		LogInUser:       &LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "CommentNew")
}
