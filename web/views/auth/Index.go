package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//LogOut lets the user out
func LogOut(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	cookie, err := request.Cookie("kRtrima")
	if err != http.ErrNoCookie {
		Logger.Println("Cookie was found with name kRtrima")
		cookie.MaxAge = -1 //delete the cookie
		http.SetCookie(writer, cookie)
		Logger.Println("Cookie has now been deleted!!")
		err := m.Sessions.Find("_id", cookie.Value)
		if err != nil {
			Logger.Println("Cannot Find session")
			http.Redirect(writer, request, "/login", 302)
			return
		}

		Logger.Println("Valid Session Was Found!!")
		if _, err := m.Sessions.DeleteItem(m.SP.ID); err != nil {
			Logger.Println("Not able to Delete the session!!")
			http.Redirect(writer, request, "/login", 302)
			return
		}
		Logger.Println("Session Was Deleted Successfully!!")
		//delete a user struct with session uuid as nil
		// update := bson.M{
		// 	"salt": "",
		// }

		//remove the user ID from the session

		Logger.Println("Login Pointer to user has been cleared!!")
		m.LIP = nil

		// if _, err := m.Users.UpdateItem(m.SP.Salt, update); err != nil {
		// 	Logger.Println("Not able to remove session ID from User!!")
		// 	http.Redirect(writer, request, "/login", 302)
		// 	return
		// }
		// Logger.Println("Session was successfully removed from user!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}
	Logger.Println("No Cookie was found with kRtrima")
	http.Redirect(writer, request, "/login", 302)
}
