package auth

import (
    "github.com/julienschmidt/httprouter"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
)

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := m.FindSession("Uuid", cookie.Value)
        m.DeleteItem(session.ID, m.Sessions)
	}
	http.Redirect(writer, request, "/", 302)
}