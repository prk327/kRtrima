package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Web() {

	p("kRtrima App", version(), "started at", config.Address)

	mux := httprouter.New()
	// handle static assets
	mux.ServeFiles("/resources/*filepath", http.Dir(config.Static))
	// index
	mux.GET("/", landingPage)
	mux.GET("/Dashboard", dashboard)
	mux.GET("/Dashboard/New", dataForm)
	mux.POST("/Dashboard", postData)
	//initializing the server
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	danger(server.ListenAndServe())
	return
}
