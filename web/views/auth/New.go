package auth

import (
    "github.com/julienschmidt/httprouter"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "login")
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}