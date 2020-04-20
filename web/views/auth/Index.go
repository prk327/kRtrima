package auth

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
	"regexp"
)

// GET /logout
// Logs the user out
func LogOut(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	cookie, err := request.Cookie("kRtrima")
	if err != http.ErrNoCookie {
		fmt.Println("Cookie was found with name kRtrima")
		cookie.MaxAge = -1 //delete the cookie
		http.SetCookie(writer, cookie)
		fmt.Println("Cookie has now been deleted!!")
		session, err := m.Findmodel("salt", cookie.Value, m.Sessions)
		if err != nil {
			fmt.Println("Cannot Find session")
			http.Redirect(writer, request, "/login", 401)
			return
		}

		re := regexp.MustCompile(`"(.*?)"`)

		rStr := fmt.Sprintf(`%v`, session[0]["_id"])

		res1 := re.FindStringSubmatch(rStr)[1]

		fmt.Println("Valid Session Was Found!!")
		if _, err := m.DeleteItem(res1, m.Sessions); err != nil {
			fmt.Println("Not able to Delete the session!!")
			http.Redirect(writer, request, "/login", 401)
			return
		}
		fmt.Println("Session Was Deleted Successfully!!")
		//delete a user struct with session uuid as nil
		update := bson.M{
			"salt": "",
		}

		rStr = fmt.Sprintf(`%v`, session[0]["userid"])

		res1 = re.FindStringSubmatch(rStr)[1]

		//remove the user ID from the session
		if _, err := m.UpdateItem(res1, update, m.Users); err != nil {
			fmt.Println("Not able to remove session ID from User!!")
			http.Redirect(writer, request, "/login", 401)
			return
		}
		fmt.Println("Session was successfully removed from user!!")
		http.Redirect(writer, request, "/Dashboard", 302)
		return
	}
	fmt.Println("No Cookie was found with kRtrima")
	http.Redirect(writer, request, "/login", 302)
}
