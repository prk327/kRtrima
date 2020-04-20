package dashboard

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	m "kRtrima/plugins/database/mongoDB/models"
	//    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"net/http"
	"regexp"
)

func Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	//	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	//	docID, err := primitive.ObjectIDFromHex(res1)
	//	if err != nil {
	//		danger(err)
	//	}

	m.DeleteItem(res1, m.Collection)

	http.Redirect(writer, request, "/Dashboard", 302)
}
