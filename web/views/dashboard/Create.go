package dashboard

import (
	"time"

	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Create is used to create a new dashboard
func Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		Logger.Println("Not able to get the form detail!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}

	newItem := m.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
		User:        LIP.ID,
		CreatedAt:   time.Now(),
	}

	_, err = m.Threads.AddItem(newItem)
	if err != nil {
		Logger.Println("Not able to add new thread to DB!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	http.Redirect(writer, request, "/Dashboard", 302)
}
