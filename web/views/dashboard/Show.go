package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"

	"github.com/julienschmidt/httprouter"

	"net/http"
)

//Show function is used to display the detail thread
func Show(w http.ResponseWriter, request *http.Request, p httprouter.Params) {

	// //get the user and assign to User UP struct
	// err := m.GetUserbyUUID("kRtrima", w, request)
	// if err != nil {
	// 	Logger.Println("Not able to find the user by UUID!!")
	// 	http.Redirect(w, request, "/login", 401)
	// 	return
	// }

	var thread m.Thread

	docit, err := m.ToDocID(p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to get the docid")
		http.Redirect(w, request, "/Dashboard", 302)
		return
	}

	//get the thread and assign to Thread TP struct
	err = m.Threads.Find("_id", docit, &thread)
	if err != nil {
		Logger.Println("Not able to Find the thread by ID!!")
		http.Redirect(w, request, "/Dashboard", 302)
		return
	}

	var up m.User

	//get the User and assign to User UP struct
	err = m.Users.Find("_id", thread.User, &up)
	if err != nil {
		Logger.Println("Not able to Find the user by ID!!")
	}

	var cmt []m.Comment

	//get the comment and assign to Comment CP struct
	err = m.Comments.FindbyKeyValue("thread", thread.ID, &cmt)
	if err != nil {
		Logger.Println("Not able to Find The Comments by ID!!")
	}

	// get the list of mongo collections
	coll, err := m.ShowCollectionNames(m.DB)
	if err != nil {
		Logger.Println("Not able to Get the list of Collection!!")
	}

	dashlist := m.FindDetails{
		CollectionNames: coll,
		ContentDetails:  &thread,
		Comments:        cmt,
		User:            &up,
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(w, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "ThreadShowContent")
}
