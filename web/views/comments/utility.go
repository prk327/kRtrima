package comments

import (
	"fmt"
	"html/template"
	"net/http"
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
