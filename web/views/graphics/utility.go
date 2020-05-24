package graphics

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//Logger is used to log all the msg in dashboard module
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("kRtrima_Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "Dashboard: ", log.Ldate|log.Ltime|log.Lshortfile)
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

// Convenience function to redirect to the error message page
func errormessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// for logging
func info(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

func danger(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Fatalln(args...)
}

func warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

// version
func version() string {
	return "0.0.1"
}