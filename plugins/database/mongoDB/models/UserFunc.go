package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	// for BSON ObjectID
)

//CreateSession a new session for an existing user
// func (user *User) CreateSession() (err error) {

// 	// Create a struct type to handle the session for login
// 	statement := Session{
// 		Salt:      user.Salt,
// 		CreatedAt: time.Now(),
// 	}

// 	_, err = Sessions.AddItem(statement)
// 	if err != nil {
// 		Logger.Fatalln(err)
// 	}

// 	return
// }

//GetUserbyUUID is used to getting the user by session uuid
// func GetUserbyUUID(cookieName string, writer http.ResponseWriter, request *http.Request) (err error) {
// 	cookie, err := request.Cookie(cookieName)
// 	if err != http.ErrNoCookie {
// 		Logger.Println("Cookie found for user by session uuid ")
// 		errr := Sessions.Find("_id", cookie.Value)
// 		if errr != nil {
// 			Logger.Println("Cannot Find session by ID")
// 			http.Redirect(writer, request, "/login", 401)
// 			return
// 		}
// 		Logger.Println("Valid Session Was Found in get user by salt")
// 		err = LogInUser.Find("salt", SP.Salt)
// 		if err != nil {
// 			Logger.Println("Cannot Find user with salt")
// 			http.Redirect(writer, request, "/login", 401)
// 			return
// 		}
// 	} else {
// 		fmt.Println("No Cookie was found with Name in get user by uuid")
// 		http.Redirect(writer, request, "/login", 401)
// 		return
// 	}
// 	return
// }

// AddToHeader add the key value to request header
func AddToHeader(key string, value interface{}, request *http.Request) (err error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		Logger.Println("Not able to add the Login Detail to Header!!")
	}

	fmt.Println(string(bytes))

	request.Header.Set(key, string(bytes))

	return
}

// GetLogInUser extract the user info from the header
func GetLogInUser(key string, result interface{}, request *http.Request) (err error) {
	bytes := []byte(request.Header.Get(key))

	// var LIU m.LogInUser
	err = json.Unmarshal(bytes, result)
	if err != nil {
		Logger.Println("Not able to find the Login User!!")
		// http.Redirect(writer, request, "/", 302)
		// return
	}

	fmt.Println(result)

	return
}
