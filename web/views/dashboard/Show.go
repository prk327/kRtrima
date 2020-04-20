package dashboard

import (
	"github.com/julienschmidt/httprouter"
	//    "go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"fmt"
	m "kRtrima/plugins/database/mongoDB/models"
	"log"
	"net/http"
	"regexp"
)

func Show(w http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	user, err := m.GetUserbyUUID("kRtrima", w, request)
	if err != nil {
		fmt.Println("Not able to find the user by UUID!!")
		log.Fatalln(err)
		http.Redirect(w, request, "/login", 401)
		return
	}

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	//	docID, err := primitive.ObjectIDFromHex(res1)
	//	if err != nil {
	//		danger(err)
	//	}

	//     fmt.Println(user)/

	dashlist := m.FindDetails{
		CollectionNames: m.ShowCollectionNames(m.DB),
		ContentDetails:  m.FindItem(res1, m.Collection),
		Comments:        m.FindAllCommentByID("thread", res1, m.Comments),
		User:            user,
	}

	if dashlist.User != nil {
		fmt.Println(dashlist.User.Name)
	}

	generateHTML(w, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "showMoreinfo")
}
