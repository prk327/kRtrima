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
		Salt:      CreateUUID(),
		Email:     user.Email,
		UserID:    user.ID,
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
		errr := Sessions.Find("salt", cookie.Value)
		if errr != nil {
			Logger.Println("Cannot Find session in getuserby uuid")
			http.Redirect(writer, request, "/login", 401)
			return
		}
		Logger.Println("Valid Session Was Found in get user by uuid")
		err = Users.Find("_id", SP.UserID)
		if err != nil {
			Logger.Println("Cannot Find user with uuid")
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
