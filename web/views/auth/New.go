package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LogIn Show the login page
func LogIn(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	dashlist := m.MainCongifDetails{
		LogInUser: m.LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "LoginContent")
}

//SignUp Show the signup page
func SignUp(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	dashlist := m.MainCongifDetails{
		LogInUser: m.LIP,
	}

	generateHTML(writer, &dashlist, "Layout", "LoginLetfSideBar", "LoginTopSidebar", "LoginModal", "Register")
}
