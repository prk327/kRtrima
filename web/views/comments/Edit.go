package comments

import (
	"github.com/julienschmidt/httprouter"

	m "kRtrima/plugins/database/mongoDB/models"

	"net/http"
)

// Edit is used to edit the thread
func Edit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	var TP m.Thread

	//find the thread by id and assign it to TP
	err := m.Threads.Find("_id", p.ByName("id"), &TP)
	if err != nil {
		Logger.Println("Not able to Find the comment by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	var CP m.Comment

	//find the comment by id and assign it to CSL
	err = m.Comments.Find("_id", p.ByName("cid"), &CP)
	if err != nil {
		Logger.Println("Not able to Find the comment by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	var UP m.User

	//get the User and assign to User UP struct
	err = m.Users.Find("_id", CP.User, &UP)
	if err != nil {
		Logger.Println("Not able to Find the user by ID!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}

	if UP.Email != LIP.Email {
		Logger.Println("You are not authorized to edit this comment!!")
		http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
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
		ContentDetails:  &TP,
		SingleComment:   &CP,
		User:            &UP,
		LogInUser:       &LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "CommentEdit")
}
