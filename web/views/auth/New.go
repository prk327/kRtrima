package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LogIn Show the login page
func LogIn(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var LIP m.LogInUser

	err := m.GetLogInUser("User", &LIP, request)
	if err == nil {
		Logger.Println("User Already Login!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}

	dashlist := m.MainCongifDetails{
		LogInUser: &LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "LoginContent")
}

//SignUp Show the signup page
func SignUp(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var LIP m.LogInUser
	err := m.GetLogInUser("User", &LIP, request)
	if err != nil {
		Logger.Printf("Failed to get the login details %v\n", err)
	}

	dashlist := m.MainCongifDetails{
		LogInUser: &LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "Register")
}
