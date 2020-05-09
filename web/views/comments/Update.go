package comments

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

//Update is used to updatde the comment
func Update(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	err := request.ParseForm()
	if err != nil {
		Logger.Println("Not able to Get the form Detail!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	update := m.Comment{
		Comment:   request.Form["comment"][0],
		CreatedAt: time.Now(),
	}

	msg, err := m.Comments.UpdateItem(p.ByName("cid"), update)
	if err != nil {
		Logger.Println("Not able to Updated the comment to DB!!")
		http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
		return
	}
	Logger.Println(msg)

	http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
}
