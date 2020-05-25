package models

import (
	"encoding/json"
	"net/http"
)

// AddToHeader add the key value to request header
func AddToHeader(key string, value interface{}, request *http.Request) (err error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		Logger.Println("Not able to add the Login Detail to Header!!")
	}

	request.Header.Set(key, string(bytes))

	return
}

// GetLogInUser extract the user info from the header
func GetLogInUser(key string, result interface{}, request *http.Request) (err error) {
	bytes := []byte(request.Header.Get(key))

	err = json.Unmarshal(bytes, result)
	if err != nil {
		Logger.Println("Not able to find the Login User!!")
		// http.Redirect(writer, request, "/", 302)
		// return
	}

	return
}
