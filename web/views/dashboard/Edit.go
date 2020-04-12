package dashboard

import (
    "github.com/julienschmidt/httprouter"
    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "regexp"
    "fmt"
)

func Edit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		danger(err)
	}
    
	dashlist := m.FindDetails{
        CollectionNames: m.ShowCollectionNames(m.DB),
        ContentDetails: m.FindItem(docID, m.Collection),
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "editData")
}