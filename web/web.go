package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Web() {
    
	//initializing the router
	mux := httprouter.New()
	// handle static assets
	mux.ServeFiles("/resources/*filepath", http.Dir(config.Static))
	// Home Page
	mux.GET("/", landingPage)
	//Display a list of all the Dashboard
	mux.GET("/Dashboard", dashboard)
	//Display form to create a dashboard
	mux.GET("/Dashboard/New", dataForm)
	//Shows the info about a dashboard
	mux.GET("/Dashboard/show/:id", detailDash)
	//Add new dashboard into showpage
	mux.POST("/Dashboard", postData)
	//initializing the server
    p("kRtrima App", version(), "started at", config.Address)
	server := http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	danger(server.ListenAndServe())
}
