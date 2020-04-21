package auth

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"golang.org/x/crypto/bcrypt"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"regexp"
)

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		http.Redirect(w, request, "/login", 400)
		return
	}
	fmt.Println("LogIN Form Parsed Successfully!!")

	user, err := m.FindUser("email", request.Form["email"][0], m.Users)
	if err != nil {
		fmt.Println("Not Able to Found a Valid User ID with this Email")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 500)
		return
	}

	if user == nil {
		fmt.Println("Invalid User Please Register!!")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 500)
		return
	}

	var authUser m.User

	mapstructure.Decode(user[0], &authUser)

	fmt.Println("Register User Found!!")

	//converting the id to primitive object id
	authUser.ID = primitive.ObjectID(user[0]["_id"].(primitive.ObjectID))
	//Converting the primitive Binary to []byte
	hashed := primitive.Binary(user[0]["hash"].(primitive.Binary)).Data

	//check for valid cookie
	_, err = request.Cookie("kRtrima") //Grab the cookie from the header
	if err == http.ErrNoCookie {
		fmt.Println("No Cookie was Found with Name kRtrima")
		//Check if the user has a valid session id
		if authUser.Salt != "" {
			fmt.Println("User already have a valid session!!")
			//Delete the old session
			if _, err := m.DeleteByID(authUser.Session, m.Sessions); err != nil {
				fmt.Println("Not able to Delete the session!!")
                http.Redirect(w, request, "/login", 401)
				return
			}
			//session was deleted

			//delete a user struct with session uuid as nil
			update := bson.M{
				"salt":    "",
				"session": nil,
			}
            
            re := regexp.MustCompile(`"(.*?)"`)

            rStr := fmt.Sprintf(`%v`, user[0]["_id"])

            res1 := re.FindStringSubmatch(rStr)[1]

			//remove the user ID from the session
			if _, err := m.UpdateItem(res1, update, m.Users); err != nil {
				fmt.Println("Not able to remove session ID from User!!")
				http.Redirect(w, request, "/login", 401)
				return
			}
			fmt.Println("Session was successfully removed from user!!")
		}
		//user dont have a valid session go to next
	} else {
		fmt.Println("Cookie was Found with Name kRtrima during login")
		// Compare if the user already has a valid session
		if authUser.Salt != "" {
			fmt.Println("User already have a valid session!!")
			//Another session already active, please logout and relogin
			http.Redirect(w, request, "/login", http.StatusUnauthorized)
			return
		}
		fmt.Println("User do not have a valid session!!")
//        reset the cookie with new info
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword(hashed, []byte(request.Form["password"][0])); err != nil {
		fmt.Println("Password Did Not Matched!!")
		// If the two passwords don't match, return a 401 status
		http.Redirect(w, request, "/login", 401)
		return
	}

	fmt.Println("Password Matched Successfully")

	//create a new session
	ssid, uuid, err := authUser.CreateSession()
	if err != nil {
		fmt.Println("Cannot Create a Valid Session!!")
		http.Redirect(w, request, "/login", 401)
		return
	}
	fmt.Println("New Session Was Created Successfully!!")

	//create a user struct with session uuid
	update := m.User{
		Salt:    uuid,
		Session: ssid,
	}

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, authUser.ID)

	res1 := re.FindStringSubmatch(rStr)[1]

	//add the new session to the user db
	if _, err := m.UpdateItem(res1, update, m.Users); err != nil {
		fmt.Println("Cannot Insert Session ID to User!!")
		http.Redirect(w, request, "/login", 401)
		return
	}
	fmt.Println("Session ID Was Inserted to User!!")

    cookie := http.Cookie{
		Name:     "kRtrima",
		Value:    uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Cookie was assigned successfully")

	fmt.Println("Authentication SuccessFul")
	http.Redirect(w, request, "/Dashboard", 302)

}

//to check if the user is logged in
func GetSession(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		cookie, err := r.Cookie("kRtrima") //Grab the cookie from the header
		if err == http.ErrNoCookie {
			fmt.Println("No Cookie was Found with Name kRtrima")
			//session is missing, returns with error code 403 Unauthorized
			http.Redirect(w, r, "/login", 302)
			//			w.WriteHeader(http.StatusForbidden)
			return
		}

		fmt.Println("Cookie was Found with Name kRtrima")

		if _, err = m.Findmodel("salt", cookie.Value, m.Sessions); err != nil {
			fmt.Println("Cannot found a valid User Session!!")
			//session is missing, returns with error code 403 Unauthorized
			http.Redirect(w, r, "/login", 302)
			return
		}
		fmt.Println("Valid User Session was Found!!")
		h(w, r, ps)
	}
}
