package auth

import (
    "github.com/julienschmidt/httprouter"
    m "kRtrima/plugins/database/mongoDB/models"
    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
    "github.com/mitchellh/mapstructure"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "regexp"
    "fmt"
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

	user, err := m.FindUser("email",request.Form["email"][0], m.Users)
	if err != nil {
        fmt.Println("Not Able to Found a Valid User ID with this Email")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 500)	
		return
	}

    if user == nil{
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
 
    // Compare if the user already has a valid session
    if authUser.Salt != "" {
        fmt.Println("User already have a valid session!!")
		//Another session already active, please logout and relogin
        http.Redirect(w, request, "/login", http.StatusUnauthorized)	
        return
	}
    
    fmt.Println("User don't have a valid session!!")
    
    // Compare the stored hashed password, with the hashed version of the password that was received
    if err = bcrypt.CompareHashAndPassword(hashed, []byte(request.Form["password"][0])); err != nil {
        fmt.Println("Password Did Not Matched!!")
		// If the two passwords don't match, return a 401 status
		http.Redirect(w, request, "/login", 401)	
        return
	}
    
    fmt.Println("Password Matched Successfully")
    
    //create a new session
    _, uuid, err := authUser.CreateSession()
    if err != nil {
			fmt.Println("Cannot Create a Valid Session!!")
            http.Redirect(w, request, "/login", 401)
            return
		}
    fmt.Println("New Session Was Created Successfully!!")
    
    //create a user struct with session uuid
    update := m.User{
        Salt: uuid,
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

//// Checks if the user is logged in and has a session, if not err is not nil
//func GetSession(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) (sess m.Session, err error) {
//	cookie, err := request.Cookie("_cookie")
//	if err == nil {
//		if _, err = m.FindSession("Salt", cookie.Value, m.Sessions); err != nil {
//			err = errors.New("Invalid session")
//            http.Redirect(writer, request, "/login", 302)
//		}
//	}
//	return
//}



//to override the post method to follow the restful convention
func GetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user has a valid session.
		cookie, err := r.Cookie("kRtrima")
	if err != http.ErrNoCookie {
        fmt.Println("Cookie was Found with Name kRtrima")
		if _, err = m.Findmodel("salt", cookie.Value, m.Sessions)
        err != nil {
			fmt.Println("Cannot found a valid User Session!!")
            http.Redirect(w, r, "/login", 302)
            return
		}
        fmt.Println("Valid User Session was Found!!")
        // Call the next handler in the chain.
		next.ServeHTTP(w, r)
	}
        fmt.Println("Cannot Found a Valid User Session!!")
		// Call the next handler in the chain.
//		next.ServeHTTP(w, r)
        http.Redirect(w, r, "/login", 302)
            return
	})
}