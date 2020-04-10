package comments

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive" // for BSON ObjectID
    m "kRtrima/plugins/database/mongoDB/models"
	"kRtrima/plugins/database/mongoDB"
    "html/template"
	"net/http"
	"regexp"
    "log"
)


// parse HTML templates
// pass in a list of file names, and get a template
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var t *template.Template
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/ui/template/comments/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(writer, "layout", data)
}


func New(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
    
    re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		log.Fatalln(err)
	}

	dashlist := m.FindDetails{
		mongoDB.ShowCollectionNames(mongoDB.DB),
		mongoDB.FindItem(docID, mongoDB.Collection),
	}
    
	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}


func Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
    
    re := regexp.MustCompile(`"(.*?)"`)

	rStr := fmt.Sprintf(`%v`, p.ByName("id"))

	res1 := re.FindStringSubmatch(rStr)[1]

	// Create a BSON ObjectID by passing string to ObjectIDFromHex() method
	docID, err := primitive.ObjectIDFromHex(res1)
	if err != nil {
		log.Fatalln(err)
	}
    
    fmt.Println(request.Form["text"][0])
    fmt.Println(request.Form["author"][0])
    
    fmt.Println(docID)
    
//	newItem := mongoDB.Commentschema{
//		Comment:        request.Form["text"][0],
//		Author:       request.Form["author"][0],
//        Thread: docID,
//	}
//
//	logger.Println(mongoDB.AddItem(newItem, mongoDB.Collection))

	http.Redirect(writer, request, "/Dashboard/show/" + p.ByName("id"), 302)

}