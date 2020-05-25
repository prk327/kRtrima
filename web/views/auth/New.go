package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LogIn Show the login page
func LogIn(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var dashlist m.MainCongifDetails

	var LIP m.LogInUser

	err := m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "LoginContent")
}

//SignUp Show the signup page
func SignUp(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var dashlist m.MainCongifDetails

	var LIP m.LogInUser

	err := m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "Register")
}
