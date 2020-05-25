package web

import (
	U "kRtrima/web/views/auth"
	C "kRtrima/web/views/comments"
	D "kRtrima/web/views/dashboard"
	G "kRtrima/web/views/graphics"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Web is to initialize the we server
func Web() {

	//initializing the router
	mux := httprouter.New()
	// handle static assets
	mux.ServeFiles("/resources/*filepath", http.Dir(config.Static))
	// Home Page
	mux.GET("/", U.IsLogIn(D.Home))
	//show register user route
	mux.GET("/register", U.IsLogIn(U.SignUp))
	//register user route
	mux.POST("/register", U.IsLogIn(U.Create))
	//show login page
	mux.GET("/login", U.IsLogIn(U.LogIn))
	//authenticate user for login
	mux.POST("/login", U.IsLogIn(U.Authenticate))
	//show logout page
	mux.GET("/logout", U.GetSession(U.LogOut))

	// Auth route
	//Display a list of all the Dashboard Index page
	mux.GET("/Dashboard", U.GetSession(D.Index))
	//Shows the info about a dashboard Show
	mux.GET("/Dashboard/show/:id", U.GetSession(D.Show))
	//Display form to create a dashboard New
	mux.GET("/Dashboard/New", U.GetSession(D.New))
	//Add new dashboard into showpage Create
	mux.POST("/Dashboard", U.GetSession(D.Create))
	//Show EDIT page
	mux.GET("/Dashboard/show/:id/edit", U.GetSession(D.Edit))
	//Updated Route
	mux.PUT("/Dashboard/show/:id", U.GetSession(D.Update))
	//Deleate Route
	mux.DELETE("/Dashboard/show/:id", U.GetSession(D.Delete))

	//new comment route
	mux.GET("/Dashboard/show/:id/comments/new", U.GetSession(C.New))
	//Add new comment to the show page
	mux.POST("/Dashboard/show/:id/comments", U.GetSession(C.Create))
	//Edit comment
	mux.GET("/Dashboard/show/:id/comments/show/:cid/edit", U.GetSession(C.Edit))
	//Updated comment
	mux.PUT("/Dashboard/show/:id/comments/show/:cid", U.GetSession(C.Update))
	//Deleate comment
	mux.DELETE("/Dashboard/show/:id/comments/show/:cid", U.GetSession(C.Delete))

	// The route for SVG Element testing
	//Display a list of all the Dashboard Index page
	mux.GET("/canvassvg", U.GetSession(G.Index))

	//initializing the server
	p("kRtrima App", version(), "started at", config.Address)
	server := http.Server{
		Addr:    config.Address,
		Handler: methodOverride(mux),
	}
	danger(server.ListenAndServe())

}
