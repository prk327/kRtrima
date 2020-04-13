package dashboard

import (
    "github.com/julienschmidt/httprouter"
//    "kRtrima/plugins/database/mongoDB"
//    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "regexp"
    "fmt"
)

func Update(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	update := m.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
	}

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
//	docID, err := primitive.ObjectIDFromHex(res1)
//	if err != nil {
//		danger(err)
//	}

	m.UpdateItem(res1, update, m.Collection)

    http.Redirect(writer, request, "/Dashboard/show/" + p.ByName("id") , 302)
}