package graphics

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Index is to show the home page
func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var dashlist m.MainCongifDetails

	_, err := request.Cookie("kRtrima") //Grab the cookie from the header
	if err == http.ErrNoCookie {
		Logger.Println("No Cookie was Found with Name kRtrima")

	} else {
		Logger.Println("Cookie was Found with Name kRtrima")
		var LIP m.LogInUser
		err := m.GetLogInUser("User", &LIP, request)
		if err != nil {
			Logger.Printf("Failed to get the login details %v\n", err)
		}

		dashlist = m.MainCongifDetails{
			LogInUser: &LIP,
		}
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "GraphicContent")
}
