package dashboard

import (
	"github.com/julienschmidt/httprouter"
	//    "kRtrima/plugins/database/mongoDB"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
)

// Create is used to create a new dashboard
func Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		Logger.Println("Not able to get the form detail!!")
		http.Redirect(writer, request, "/home", 401)
		return
	}

	newItem := m.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
	}

	_, err = m.Threads.AddItem(newItem)
	if err != nil {
		Logger.Println("Not able to add new thread to DB!!")
		http.Redirect(writer, request, "/home", 401)
		return
	}

	http.Redirect(writer, request, "/Dashboard", 302)
}
