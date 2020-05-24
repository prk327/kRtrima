package comments

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Create is used to create a new msg
func Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		Logger.Println("Not able to get The Form Detail!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := m.ToDocID(p.ByName("id"))
	if err != nil {
		Logger.Println("Not able to get the id of the Thread!!")
		http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
		return
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}
	newItem := m.Comment{
		Comment:   request.Form["comment"][0],
		Author:    LIP.Name,
		Thread:    docID,
		User:      LIP.ID,
		CreatedAt: time.Now(),
	}

	_, err = m.Comments.AddItem(newItem)
	if err != nil {
		Logger.Println("Not able to add new Comment to DB!!")
		http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
		return
	}

	http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)

}
