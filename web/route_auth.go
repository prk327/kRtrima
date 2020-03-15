package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//dummy data for testing
var p1 = map[string]string{
	"Name":  "Circuit",
	"Image": "/resources/images/circuit.jpg",
}
var p2 = map[string]string{
	"Name":  "City At Night",
	"Image": "/resources/images/CityNight.jpg",
}
var p3 = map[string]string{
	"Name":  "Pyramid",
	"Image": "/resources/images/piramid.jpg",
}
var dashList = []map[string]string{p1, p2, p3}

func dashboard(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, dashList, "layout", "leftsidebar", "topsidebar", "modal", "dashboard")
}

func postData(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		danger(err)
	}
	name := request.Form["name"]
	image := request.Form["image"]
	p4 := map[string]string{
		"Name":  name[0],
		"Image": image[0],
	}
	dashList = append(dashList, p4)
	generateHTML(writer, dashList, "layout", "leftsidebar", "topsidebar", "modal", "dashboard")
}

func landingPage(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, "This is the Landing Page", "landing")
}

func dataForm(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	generateHTML(writer, nil, "layout", "leftsidebar", "topsidebar", "modal", "newDForm")
}
