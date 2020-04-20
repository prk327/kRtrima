package auth

import (
	"github.com/julienschmidt/httprouter"
	//    m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
)

// GET /login
// Show the login page
func LogIn(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "login_new")
}

// GET /signup
// Show the signup page
func SignUp(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "signup")
}
