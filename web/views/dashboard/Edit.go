package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Edit is used to edit the thread
func Edit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	var TP m.Thread

	//find the thread by id and assign it to TP
	err := m.Threads.Find("_id", p.ByName("id"), &TP)
	if err != nil {
		Logger.Println("Not able to Find the thread by ID!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	var UP m.User

	//get the User and assign to User UP struct
	err = m.Users.Find("_id", TP.User, &UP)
	if err != nil {
		Logger.Println("Not able to Find the user by ID!!")
		http.Redirect(writer, request, "/", 302)
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
		ContentDetails:  &TP,
		User:            &UP,
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "ThreadEdit")
}
