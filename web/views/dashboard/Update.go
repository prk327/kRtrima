package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Update is used to updatde the thraed
func Update(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	err := request.ParseForm()
	if err != nil {
		Logger.Println("Not able to Get the form data!!")
		http.Redirect(writer, request, "/", 302)
		return
	}

	update := m.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
		User:        m.LIP.ID,
	}

	msg, err := m.Threads.UpdateItem(p.ByName("id"), update)
	if err != nil {
		Logger.Println("Not able to Updated the thread to DB!!")
		http.Redirect(writer, request, "/", 302)
		return
	}
	Logger.Println(msg)

	http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
}
