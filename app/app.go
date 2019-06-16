package app

import (
	"log"
	"net/http"

	. "github.com/ConstantineGochev/go_api/config"
	"github.com/ConstantineGochev/go_api/controllers"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

type App struct {
	Router   *mux.Router
	db       *mgo.Database
	Server   string
	Database string
}

var config = Config{}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize() {
	config.Read()
	a.Database = config.Database
	a.Server = config.Server
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	a.db = session.DB(a.Database)
	a.Router = mux.NewRouter()
	a.SetRouters()
}

// SetRouters sets the all required routers
func (a *App) SetRouters() {
	// Routing for handling the news
	a.Get("/news", a.handle_request(controllers.GetAllNews))
	a.Post("/news", a.handle_request(controllers.CreateNewNews))
	a.Get("/", a.handle_request(controllers.Test))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestControllersFunction func(db *mgo.Database, w http.ResponseWriter, r *http.Request)

func (a *App) handle_request(controllers RequestControllersFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		controllers(a.db, w, r)
	}
}
