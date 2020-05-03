package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LogIn Show the login page
func LogIn(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "login_new")
}

//SignUp Show the signup page
func SignUp(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "signup")
}
