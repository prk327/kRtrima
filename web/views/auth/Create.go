package auth

import (
    "github.com/julienschmidt/httprouter"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "time"
)

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
    
	user := m.User{
        Uuid:   m.CreateUUID(),
        Email:    request.PostFormValue("email"),
		Name:     request.PostFormValue("name"),
		Password:  m.Encrypt(request.PostFormValue("password")),
        CreatedAt: time.Now(),
	}
    
    _, err := m.AddItem(user, m.Users)
    if err != nil {
		danger(err, "Cannot create user")
	}
    
	http.Redirect(writer, request, "/login", 302)
}