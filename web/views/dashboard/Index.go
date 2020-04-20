package dashboard

import (
	"github.com/julienschmidt/httprouter"

	//    "kRtrima/plugins/database/mongoDB"
	"fmt"
	m "kRtrima/plugins/database/mongoDB/models"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	user, err := m.GetUserbyUUID("kRtrima", writer, request)
	if err != nil {
		fmt.Println("Not able to find the user by user in Index")
		//            http.Redirect(writer, request, "/login", 401)
		//            return
	}
	fmt.Printf("%T", user)

	dashlist := m.MainCongifDetails{
		CollectionNames: m.ShowCollectionNames(m.DB),
		ContentDetails:  m.FindAllItem(10, m.Collection),
		User:            user,
	}

	//    if dashlist.User != nil{
	//        fmt.Println(dashlist.User.Name)
	//    }
	//

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "index")
}
