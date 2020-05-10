package auth

import (
	"fmt"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//LogOut lets the user out
func LogOut(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	cookie, err := request.Cookie("kRtrima")
	if err != http.ErrNoCookie {
		Logger.Println("Cookie was found with name kRtrima")
		cookie.MaxAge = -1 //delete the cookie
		http.SetCookie(writer, cookie)
		Logger.Println("Cookie has now been deleted!!")

		// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
		docID, err := primitive.ObjectIDFromHex(cookie.Value)
		if err != nil {
			fmt.Printf("Cannot Convert %T type to object id", cookie.Value)
			Logger.Println(err)
		}

		if err = m.Sessions.Find("_id", docID); err != nil {
			Logger.Println("Cannot found a valid User Session!!")
			//session is missing, returns with error code 403 Unauthorized
			http.Redirect(writer, request, "/login", 302)
			return
		}

		Logger.Println("Valid User Session was Found!!")

		if _, err := m.Sessions.DeleteItem(m.SP.ID); err != nil {
			Logger.Println("Not able to Delete the session!!")
			http.Redirect(writer, request, "/login", 302)
			return
		}

		// reset the SP
		m.SP = nil
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
