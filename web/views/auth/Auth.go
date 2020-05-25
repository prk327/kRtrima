package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter" // for BSON ObjectID
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//Authenticate the user given the email and password
func Authenticate(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		Logger.Println("Not able to parse the form!!")
		//remove the user ID from the session
		request.Header.Del("User")
		http.Redirect(w, request, "/login", 302)
		return
	}
	Logger.Println("LogIn Form Parsed Successfully!!")

	var UP m.User

	err = m.Users.Find("email", request.Form["email"][0], &UP)
	if err != nil {
		Logger.Printf("Not Able to Found a Valid User ID with Email: %v", request.Form["email"][0])
		// If there is an issue with the database, return a 500 error
		//remove the user ID from the session
		request.Header.Del("User")
		http.Redirect(w, request, "/register", 302)
		return
	}

	if UP.Email != request.Form["email"][0] {
		Logger.Println("Invalid User Please Register!!")
		// If there is an issue with the database, return a 500 error
		//remove the user ID from the session
		request.Header.Del("User")
		http.Redirect(w, request, "/register", 302)
		return
	}

	Logger.Printf("Register User with Email %v was Found Successfully", UP.Email)

	//check for valid cookie
	// _, err = request.Cookie("kRtrima") //Grab the cookie from the header
	// if err == http.ErrNoCookie {
	// 	Logger.Println("No Cookie was Found with Name kRtrima")
	// 	//Check if the user has a valid session id
	// 	if m.UP.Salt != "" {
	// 		Logger.Println("User already have a valid session!!")
	// 		//Delete the old session
	// 		if _, err := m.Sessions.DeleteItem(m.UP.Session); err != nil {
	// 			Logger.Println("Not able to Delete the session!!")
	// 			http.Redirect(w, request, "/login", 302)
	// 			return
	// 		}
	// 		//session was deleted

	// 		//delete a user struct with session uuid as nil
	// 		update := bson.M{
	// 			"salt":    "",
	// 			"session": nil,
	// 		}

	// 		//remove the user ID from the session
	// 		if _, err := m.Users.UpdateItem(m.UP.ID, update); err != nil {
	// 			Logger.Println("Not able to remove session ID from User!!")
	// 			http.Redirect(w, request, "/login", 302)
	// 			return
	// 		}
	// 		Logger.Println("Session was successfully removed from user!!")
	// 	}
	// 	//user dont have a valid session go to next
	// } else {
	// 	Logger.Println("Cookie was Found with Name kRtrima during login")
	// 	// Compare if the user already has a valid session
	// 	if m.UP.Salt != "" {
	// 		Logger.Println("User already have a valid session!!")
	// 		//Another session already active, please logout and relogin
	// 		http.Redirect(w, request, "/Dashboard", 302)
	// 		return
	// 	}
	// 	Logger.Println("User do not have a valid session!!")
	// 	//        reset the cookie with new info
	// }

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword(UP.Hash, []byte(request.Form["password"][0])); err != nil {
		Logger.Println("Password Did Not Matched!!")
		// If the two passwords don't match, return a 401 status
		//remove the user ID from the session
		request.Header.Del("User")
		http.Redirect(w, request, "/login", 302)
		return
	}

	Logger.Println("Password Matched Successfully")

	var SSL []m.Session
	// Find all the session with user salt
	err = m.Sessions.FindbyKeyValue("salt", UP.Salt, &SSL)
	if err != nil {
		Logger.Printf("Cannot find a Valid Session for User %v", UP.Name)
		// http.Redirect(w, request, "/login", 302)
	}

	if SSL != nil {
		for _, s := range SSL {
			// 		//Delete the old session
			if _, err := m.Sessions.DeleteItem(s.ID); err != nil {
				Logger.Printf("Not able to Delete the session with ID: %v", s.ID)
				//remove the user ID from the session
				request.Header.Del("User")
				http.Redirect(w, request, "/login", 302)
				return
			}
		}
	}

	//create a new session
	// Create a struct type to handle the session for login
	statement := m.Session{
		Salt:      UP.Salt,
		CreatedAt: time.Now(),
	}

	SSID, err := m.Sessions.AddItem(statement)
	if err != nil {
		Logger.Printf("Cannot Create a Valid Session for User: %v", UP.Name)
		//remove the user ID from the session
		request.Header.Del("User")
		http.Redirect(w, request, "/login", 302)
		return
	}
	Logger.Printf("New Session Was Created Successfully for User: %v", UP.Name)

	// //create a user struct with session uuid
	// update := m.User{
	// 	Salt:    uuid,
	// 	Session: ssid,
	// }

	// //add the new session to the user db
	// if _, err := m.Users.UpdateItem(m.UP.ID, update); err != nil {
	// 	Logger.Println("Cannot Insert Session ID to User!!")
	// 	http.Redirect(w, request, "/login", 302)
	// 	return
	// }
	// Logger.Println("Session ID Was Inserted to User!!")

	// re := regexp.MustCompile(`"(.*?)"`)
	// rStr := fmt.Sprintf(`%v`, SSID.Hex())
	// rStr1 := re.FindStringSubmatch(rStr)[1]

	// fmt.Println(SSID.Hex())

	cookie := http.Cookie{
		Name:     "kRtrima",
		Value:    SSID.Hex(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	Logger.Printf("Cookie was assigned successfully for User %v", UP.Name)

	Logger.Println("Authentication SuccessFul")
	http.Redirect(w, request, "/Dashboard", 302)

}

//GetSession to check if the user has the authority to view the page
func GetSession(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		cookie, err := r.Cookie("kRtrima") //Grab the cookie from the header
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				Logger.Println("No Cookie was Found with Name kRtrima")
				//remove the user ID from the session
				r.Header.Del("User")
				http.Redirect(w, r, "/login", 302)
				return
			default:
				Logger.Println("No Cookie was Found with Name kRtrima")
				//remove the user ID from the session
				r.Header.Del("User")
				http.Redirect(w, r, "/login", 302)
				return
			}
		}

		Logger.Println("Cookie was Found with Name kRtrima")

		// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
		docID, err := primitive.ObjectIDFromHex(cookie.Value)
		if err != nil {
			Logger.Printf("Cannot Convert %T type to object id", cookie.Value)
			Logger.Println(err)
		}

		var SP m.Session
		if err = m.Sessions.Find("_id", docID, &SP); err != nil {
			Logger.Println("Cannot found a valid User Session!!")
			//session is missing, returns with error code 403 Unauthorized
			//remove the user ID from the session
			r.Header.Del("User")
			http.Redirect(w, r, "/login", 302)
			return
		}

		Logger.Println("Valid User Session was Found!!")

		var UP m.LogInUser

		err = m.Users.Find("salt", SP.Salt, &UP)
		if err != nil {
			Logger.Println("Cannot Find user with salt")
			//Delete the old session
			if _, err := m.Sessions.DeleteItem(SP.ID); err != nil {
				Logger.Printf("Not able to Delete the session with ID: %v", SP.ID)
				//remove the user ID from the session
				r.Header.Del("User")
				http.Redirect(w, r, "/login", 302)
				return
			}
			// reset session and login user
			//remove the user ID from the session
			r.Header.Del("User")
			http.Redirect(w, r, "/register", 302)
			return
		}

		var LIP m.LogInUser

		err = m.GetLogInUser("User", &LIP, r)
		if err != nil {
			m.AddToHeader("User", UP, r)
		} else if UP.Email != LIP.Email {
			//remove the user ID from the session
			r.Header.Del("User")
			m.AddToHeader("User", UP, r)
		}

		h(w, r, ps)
	}
}

//IsLogIn to check if the user is logged in
func IsLogIn(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		cookie, err := r.Cookie("kRtrima") //Grab the cookie from the header
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				Logger.Println("No Cookie was Found with Name kRtrima")
				h(w, r, ps)
				return
			default:
				Logger.Println("No Cookie was Found with Name kRtrima")
				h(w, r, ps)
				return
			}
		}

		Logger.Println("Cookie was Found with Name kRtrima")

		// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
		docID, err := primitive.ObjectIDFromHex(cookie.Value)
		if err != nil {
			Logger.Printf("Cannot Convert %T type to object id", cookie.Value)
			Logger.Println(err)
			h(w, r, ps)
			return
		}

		var SP m.Session
		if err = m.Sessions.Find("_id", docID, &SP); err != nil {
			Logger.Println("Cannot found a valid User Session!!")
			h(w, r, ps)
			return
			//session is missing, returns with error code 403 Unauthorized
		}

		Logger.Println("Valid User Session was Found!!")

		var UP m.LogInUser

		err = m.Users.Find("salt", SP.Salt, &UP)
		if err != nil {
			Logger.Println("Cannot Find user with salt")
			h(w, r, ps)
			return
		}

		var LIP m.LogInUser

		err = m.GetLogInUser("User", &LIP, r)
		if err != nil {
			m.AddToHeader("User", UP, r)
			h(w, r, ps)
			return
		} else if UP.Email != LIP.Email {
			//remove the user ID from the session
			r.Header.Del("User")
			m.AddToHeader("User", UP, r)
			h(w, r, ps)
			return
		}

		h(w, r, ps)
	}
}
