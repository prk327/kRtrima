package dashboard

import (
    "github.com/julienschmidt/httprouter"

//    "kRtrima/plugins/database/mongoDB"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
//    "fmt"
)



func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	dashlist := m.MainCongifDetails{
		m.ShowCollectionNames(m.DB),
		m.FindAllItem(10, m.Collection),
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "index")
}