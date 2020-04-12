package dashboard

import (
    "github.com/julienschmidt/httprouter"
//    "kRtrima/plugins/database/mongoDB"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
)


func Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	newItem := m.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
	}

	m.AddItem(newItem, m.Collection)

	http.Redirect(writer, request, "/Dashboard", 302)
}