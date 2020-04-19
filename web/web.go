package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
    D "kRtrima/web/views/dashboard"
    C "kRtrima/web/views/comments"
    U "kRtrima/web/views/auth"
)

func Web() {

	//initializing the router
	mux := httprouter.New()
	// handle static assets
	mux.ServeFiles("/resources/*filepath", http.Dir(config.Static))
	// Home Page
	mux.GET("/", D.Home)
    //show register user route
    mux.GET("/register", U.SignUp)
    //register user route
    mux.POST("/register", U.Create)
    //show login page
    mux.GET("/login", U.LogIn)
    //authenticate user for login
    mux.POST("/login", U.Authenticate)
    //show logout page
    mux.GET("/logout", U.LogOut)
    
    
    // Auth route
	//Display a list of all the Dashboard Index page
    mux.GET("/Dashboard", D.Index)
	//Display form to create a dashboard New
    mux.GET("/Dashboard/New", U.GetSession(D.New))
	//Add new dashboard into showpage Create
    mux.POST("/Dashboard", U.GetSession(D.Create))
	//Shows the info about a dashboard Show
    mux.GET("/Dashboard/show/:id", U.GetSession(D.Show))
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

	//initializing the server
	p("kRtrima App", version(), "started at", config.Address)
	server := http.Server{
		Addr:    config.Address,
        Handler: methodOverride(mux),
	}
	danger(server.ListenAndServe())
    
}
