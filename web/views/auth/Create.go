package auth

import (
	"fmt"
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
		http.Redirect(w, request, "/register", 302)
		return
	}
	Logger.Println("SignUp Form Parsed Successfully!!")

	var UP m.User
	//check for existing user
	err = m.Users.Find("email", request.Form["email"][0], &UP)
	if err != nil && fmt.Sprintf("%v", err) != "mongo: no documents in result" {
		Logger.Printf("Got some unexpected error %v", err)
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 302)
		return
	}

	if UP.Email == request.Form["email"][0] {
		Logger.Println("User Already Registered!!")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/login", 302)
		return
	}

	hashed, err := m.Encrypt(request.Form["password"][0])
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		Logger.Println("Soething went wrong!!")
		http.Redirect(w, request, "/login", 302)
		return
	}
	Logger.Println("New User Password was Hashed Successfully!!")

	var admin bool
	if request.Form["adminCode"][0] == "161812" {
		admin = true
	} else {
		admin = false
	}

	user := m.User{
		Salt:      m.CreateUUID(),
		Email:     request.Form["email"][0],
		Name:      request.Form["name"][0],
		Hash:      hashed,
		CreatedAt: time.Now(),
		IsAdmin:   admin,
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
