package dashboard

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func New(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}
