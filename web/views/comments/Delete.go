package comments

import (
	m "kRtrima/plugins/database/mongoDB/models"

	"github.com/julienschmidt/httprouter"

	"net/http"
)

// Delete is used to delete a particular item
func Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	//delete the comment by id
	_, err := m.Comments.DeleteItem(p.ByName("cid"))
	if err != nil {
		Logger.Println("Not able to delete comment from DB!!")
		http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
		return
	}

	http.Redirect(writer, request, "/Dashboard/show/"+p.ByName("id"), 302)
}
