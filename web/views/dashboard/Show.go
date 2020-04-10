package dashboard

import (
    "github.com/julienschmidt/httprouter"
    "kRtrima/plugins/database/mongoDB"
    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "regexp"
    "fmt"
)


func Show(w http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		danger(err)
	}

	dashlist := m.FindDetails{
		mongoDB.ShowCollectionNames(mongoDB.DB),
		mongoDB.FindItem(docID, mongoDB.Collection),
	}

	generateHTML(w, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "showMoreinfo")
}