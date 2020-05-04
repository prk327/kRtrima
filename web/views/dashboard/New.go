package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New function is used to display the new Thread form
func New(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	dashlist := m.MainCongifDetails{
		LogInUser: m.LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "ThreadNew")
}
