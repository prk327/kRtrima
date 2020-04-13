package auth

import (
    "github.com/julienschmidt/httprouter"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "golang.org/x/crypto/bcrypt"
)

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	user, err := m.FindUser("email",request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
    
    err := bcrypt.CompareHashAndPassword(user.Password, request.PostFormValue("password"))
    if err != nil{
        http.Redirect(writer, request, "/login", 302)
    }
    session, err := user.CreateSession()
    if err != nil {
			danger(err, "Cannot create session")
		}
    cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)

}

// Checks if the user is logged in and has a session, if not err is not nil
func GetSession(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) (sess m.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		if _, err = m.FindSession("Uuid", cookie.Value, m.Sessions); err != nil {
			err = errors.New("Invalid session")
		}
	}
	return
}