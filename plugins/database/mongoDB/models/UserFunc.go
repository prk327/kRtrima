package models

import (
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
)

//CreateSession a new session for an existing user
func (user *User) CreateSession() (ssid primitive.ObjectID, uuid string, err error) {

	// Create a struct type to handle the session for login
	statement := Session{
		Salt:      user.Salt,
		CreatedAt: time.Now(),
	}

	ssid, err = Sessions.AddItem(statement)
	if err != nil {
		Logger.Fatalln(err)
	}
	uuid = statement.Salt

	return
}

//GetUserbyUUID is used to getting the user by session uuid
func GetUserbyUUID(cookieName string, writer http.ResponseWriter, request *http.Request) (err error) {
	cookie, err := request.Cookie(cookieName)
	if err != http.ErrNoCookie {
		Logger.Println("Cookie found for user by session uuid ")
		errr := Sessions.Find("_id", cookie.Value)
		if errr != nil {
			Logger.Println("Cannot Find session by ID")
			http.Redirect(writer, request, "/login", 401)
			return
		}
		Logger.Println("Valid Session Was Found in get user by salt")
		err = LogInUser.Find("salt", SP.Salt)
		if err != nil {
			Logger.Println("Cannot Find user with salt")
			http.Redirect(writer, request, "/login", 401)
			return
		}
	} else {
		fmt.Println("No Cookie was found with Name in get user by uuid")
		http.Redirect(writer, request, "/login", 401)
		return
	}
	return
}
