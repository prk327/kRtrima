package comments

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
//	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	m "kRtrima/plugins/database/mongoDB/models"
	//	"kRtrima/plugins/database/mongoDB"
	//    "html/template"
//	"log"
	"net/http"
	"regexp"
)

func New(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

//	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
//	docID, err := primitive.ObjectIDFromHex(res1)
//	if err != nil {
//		log.Fatalln(err)
//	}

	dashlist := m.FindDetails{
		CollectionNames: m.ShowCollectionNames(m.DB),
		ContentDetails:  m.FindItem(res1, m.Collection),
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}
