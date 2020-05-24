package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New function is used to display the new Thread form
func New(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var LIP m.LogInUser

	err := m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}

	dashlist := m.MainCongifDetails{
		LogInUser: &LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "ThreadNew")
}
