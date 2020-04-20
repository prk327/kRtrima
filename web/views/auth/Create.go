package auth

import (
	"github.com/julienschmidt/httprouter"
	//    "github.com/satori/go.uuid"
	"fmt"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"time"
)

// POST /signup
// Create the user account
func Create(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("SignUp Form Parsed Successfully!!")

	//    check for existing user
	foundUser, err := m.FindUser("email", request.Form["email"][0], m.Users)
	if err != nil {
		fmt.Printf("Got some unexpected error")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/register", 500)
		return
	}

	if foundUser != nil {
		fmt.Printf("User Already Registered!!")
		// If there is an issue with the database, return a 500 error
		http.Redirect(w, request, "/login", 500)
		return
	}

	hashed, err := m.Encrypt(request.Form["password"][0])
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("New User Password was Hashed Successfully!!")

	user := m.User{
		Email:     request.Form["email"][0],
		Name:      request.Form["name"][0],
		Hash:      hashed,
		CreatedAt: time.Now(),
	}

	_, err = m.AddItem(user, m.Users)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("User %v Was successfully Created!!", request.Form["name"][0])

	http.Redirect(w, request, "/login", 302)
}
