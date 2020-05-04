package auth

import (
	"fmt"

	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson" // for BSON ObjectID
	"golang.org/x/crypto/bcrypt"
)

//Authenticate the user given the email and password
func Authenticate(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		Logger.Println("Not able to parse the form!!")
		http.Redirect(w, request, "/login", 302)
		return
	}
	fmt.Println("LogIN Form Parsed Successfully!!")

	err = m.Users.Find("email", request.Form["email"][0])
	if err != nil {
		Logger.Println("Not Able to Found a Valid User ID with this Email")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 302)
		return
	}

	if m.UP.Email != request.Form["email"][0] {
		Logger.Println("Invalid User Please Register!!")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 302)
		return
	}

	Logger.Println("Register User Found!!")

	//check for valid cookie
	_, err = request.Cookie("kRtrima") //Grab the cookie from the header
	if err == http.ErrNoCookie {
		Logger.Println("No Cookie was Found with Name kRtrima")
		//Check if the user has a valid session id
		if m.UP.Salt != "" {
			Logger.Println("User already have a valid session!!")
			//Delete the old session
			if _, err := m.Sessions.DeleteItem(m.UP.Session); err != nil {
				Logger.Println("Not able to Delete the session!!")
				http.Redirect(w, request, "/login", 302)
				return
			}
			//session was deleted

			//delete a user struct with session uuid as nil
			update := bson.M{
				"salt":    "",
				"session": nil,
			}

			//remove the user ID from the session
			if _, err := m.Users.UpdateItem(m.UP.ID, update); err != nil {
				Logger.Println("Not able to remove session ID from User!!")
				http.Redirect(w, request, "/login", 302)
				return
			}
			Logger.Println("Session was successfully removed from user!!")
		}
		//user dont have a valid session go to next
	} else {
		Logger.Println("Cookie was Found with Name kRtrima during login")
		// Compare if the user already has a valid session
		if m.UP.Salt != "" {
			Logger.Println("User already have a valid session!!")
			//Another session already active, please logout and relogin
			http.Redirect(w, request, "/Dashboard", 302)
			return
		}
		Logger.Println("User do not have a valid session!!")
		//        reset the cookie with new info
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword(m.UP.Hash, []byte(request.Form["password"][0])); err != nil {
		Logger.Println("Password Did Not Matched!!")
		// If the two passwords don't match, return a 401 status
		http.Redirect(w, request, "/login", 302)
		return
	}

	Logger.Println("Password Matched Successfully")

	//create a new session
	ssid, uuid, err := m.UP.CreateSession()
	if err != nil {
		Logger.Println("Cannot Create a Valid Session!!")
		http.Redirect(w, request, "/login", 302)
		return
	}
	Logger.Println("New Session Was Created Successfully!!")

	//create a user struct with session uuid
	update := m.User{
		Salt:    uuid,
		Session: ssid,
	}

	//add the new session to the user db
	if _, err := m.Users.UpdateItem(m.UP.ID, update); err != nil {
		Logger.Println("Cannot Insert Session ID to User!!")
		http.Redirect(w, request, "/login", 302)
		return
	}
	Logger.Println("Session ID Was Inserted to User!!")

	cookie := http.Cookie{
		Name:     "kRtrima",
		Value:    uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	Logger.Println("Cookie was assigned successfully")

	Logger.Println("Authentication SuccessFul")
	http.Redirect(w, request, "/Dashboard", 302)

}

//GetSession to check if the user is logged in
func GetSession(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		cookie, err := r.Cookie("kRtrima") //Grab the cookie from the header
		if err == http.ErrNoCookie {
			Logger.Println("No Cookie was Found with Name kRtrima")
			//session is missing, returns with error code 403 Unauthorized
			http.Redirect(w, r, "/login", 302)
			return
		}

		Logger.Println("Cookie was Found with Name kRtrima")

		if err = m.Sessions.Find("salt", cookie.Value); err != nil {
			Logger.Println("Cannot found a valid User Session!!")
			//session is missing, returns with error code 403 Unauthorized
			http.Redirect(w, r, "/login", 302)
			return
		}

		Logger.Println("Valid User Session was Found!!")

		err = m.LogInUser.Find("_id", m.SP.UserID)
		if err != nil {
			Logger.Println("Cannot Find user with uuid")
		}

		h(w, r, ps)
	}
}
