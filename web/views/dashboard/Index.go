package dashboard

import (
    "github.com/julienschmidt/httprouter"
    "html/template"
    "kRtrima/plugins/database/mongoDB"
    m "kRtrima/plugins/database/mongoDB/models"
    "net/http"
    "fmt"
)


// parse HTML templates
// pass in a list of file names, and get a template
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {

	var t *template.Template
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/ui/template/datasource/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(writer, "layout", data)
}


func Home(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func Index(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	dashlist := m.MainCongifDetails{
		mongoDB.ShowCollectionNames(mongoDB.DB),
		mongoDB.FindAllItem(10, mongoDB.Collection),
	}

	generateHTML(writer, &dashlist, "layout", "leftsidebar", "topsidebar", "modal", "index")
}