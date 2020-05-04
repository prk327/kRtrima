package comments

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Logger is used to log all the comments msg
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("kRtrima_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "Comment: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// parse HTML templates
// pass in a list of file names, and get a template
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var t *template.Template
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("web/ui/template/HTMLLayouts/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(writer, "layout", data)
}
