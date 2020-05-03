package auth

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

//Create is the POST route to create the user account
func Create(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		Logger.Printf("Error while parsing the Form!!")
		http.Redirect(w, request, "/register", 500)
		return
	}
	Logger.Println("SignUp Form Parsed Successfully!!")

	//check for existing user
	err = m.Users.Find("email", request.Form["email"][0])
	if err != nil {
		Logger.Printf("Got some unexpected error")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 500)
		return
	}

	if m.UP.Email == request.Form["email"][0] {
		Logger.Println("User Already Registered!!")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/login", 500)
		return
	}

	hashed, err := m.Encrypt(request.Form["password"][0])
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		Logger.Println("Soething went wrong!!")
		http.Redirect(w, request, "/login", 400)
		return
	}
	Logger.Println("New User Password was Hashed Successfully!!")

	user := m.User{
		Email:     request.Form["email"][0],
		Name:      request.Form["name"][0],
		Hash:      hashed,
		CreatedAt: time.Now(),
	}

	_, err = m.Users.AddItem(user)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	Logger.Printf("User %v Was successfully Created!!", request.Form["name"][0])

	http.Redirect(w, request, "/login", 302)
}
