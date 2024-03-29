package dashboard

import (
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegEx represents a regular expression. The Options field may contain
// individual characters defining the way in which the pattern should be
// applied, and must be sorted. Valid options as of this writing are 'i' for
// case insensitive matching, 'm' for multi-line matching, 'x' for verbose
// mode, 'l' to make \w, \W, and similar be locale-dependent, 's' for dot-all
// mode (a '.' matches everything), and 'u' to make \w, \W, and similar match
// unicode. The value of the Options parameter is not verified before being
// marshaled into the BSON format.
type RegEx struct {
	Pattern string
	Options string
}

// Home is to show the home page
func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var dashlist m.MainCongifDetails

	var LIP m.LogInUser

	err := m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(writer, &dashlist, "Landing", "LoginTopSidebar", "LandingContent")
}

//Index is used to show the threads
func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	var TSL []m.Thread
	//get the thread and assign it to slice of thread TSL
	err := m.Threads.FindAll(100, &TSL)
	if err != nil {
		Logger.Println("Not able to Find the list of Thread!!")
		http.Redirect(writer, request, "/home", 302)
		return
	}

	// get the list of mongo collections
	coll, err := m.ShowCollectionNames(m.DB)
	if err != nil {
		Logger.Println("Not able to Get the list of Collection!!")
		http.Redirect(writer, request, "/", 302)
		return
	}

	dashlist := m.MainCongifDetails{
		CollectionNames: coll,
		ContentDetails:  TSL,
	}

	var LIP m.LogInUser

	err = m.GetLogInUser("User", &LIP, request)
	if err != nil {
		dashlist.LogInUser = nil
		Logger.Printf("Failed to get the login details %v\n", err)
	} else {
		dashlist.LogInUser = &LIP
	}

	generateHTML(writer, &dashlist, "Layout", "ThreadLeftSideBar", "ThreadTopSideBar", "ThreadModal", "ThreadIndexContent")
}
