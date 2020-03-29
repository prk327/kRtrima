package web

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
	"kRtrima/plugins/database/mongoDB"
	"net/http"
	"regexp"
)

var lm = &mongoDB.Msg

type mainCongifDetails struct {
	CollectionNames []string
	ContentDetails  []*mongoDB.Thread
}

type findDetails struct {
	CollectionNames []string
	ContentDetails  *mongoDB.Thread
}

type DBConfig struct {
	Host       string
	Database   string
	Collection string
}


func dashboard(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {


	dashlist := mainCongifDetails{
		mongoDB.ShowCollectionNames(mongoDB.DB),
		mongoDB.FindAllItem(10, mongoDB.Collection),
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "index")
}

func postData(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	newItem := mongoDB.Thread{
		Name:        request.Form["name"][0],
		Image:       request.Form["image"][0],
		Description: request.Form["desc"][0],
	}
    
    logger.Println(mongoDB.AddItem(newItem, mongoDB.Collection))
    
	http.Redirect(writer, request, "/Dashboard", 302)

}

func landingPage(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func dataForm(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}

func connectDB(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}

	DB_Config := DBConfig{
		request.Form["mongo_HostName"][0],
		request.Form["mongo_DBName"][0],
		request.Form["mongo_CollName"][0],
	}

	_, mongoDB.DB = mongoDB.Connect_mongoDB(DB_Config.Host, DB_Config.Database)


	_, mongoDB.Collection = mongoDB.Cnt_Collection(DB_Config.Collection, mongoDB.DB)

	
	http.Redirect(writer, request, "/", 302)

}

func detailDash(w http.ResponseWriter, request *http.Request, p httprouter.Params) {

	re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		danger(err)
	}

	dashlist := findDetails{
		mongoDB.ShowCollectionNames(mongoDB.DB),
		mongoDB.FindItem(docID, mongoDB.Collection),
	}

	generateHTML(w, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "showMoreinfo")
}
