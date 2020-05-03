package dashboard

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New function is used to display the new Thread form
func New(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}
