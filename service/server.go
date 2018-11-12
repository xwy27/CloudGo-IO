package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	xrouter "github.com/xwy27/CloudGo-IO/router"
)

// NewServer sets the configurations of a server and returns it
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	router := mux.NewRouter()
	initRouter(router, formatter)

	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

// initial router with a given render
func initRouter(router *mux.Router, render *render.Render) {
	staticRoot, err := os.Getwd()
	if err != nil {
		panic("Could not retrive static file directory")
	}

	// path router
	router.HandleFunc("/", xrouter.HomeHandler(render))
	router.HandleFunc("/home", xrouter.HomeHandler(render))
	router.HandleFunc("/login", xrouter.LoginHandler(render))
	router.HandleFunc("/index", xrouter.LoginHandler(render))

	// Api router
	router.HandleFunc("/api/info", xrouter.InfoHandler(render))

	// Static file router
	router.PathPrefix("/templates").Handler(
		http.StripPrefix("/templates/", http.FileServer(http.Dir(staticRoot+"/templates/"))))
	router.PathPrefix("/static").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(staticRoot+"/static"))))

	// Not implement router
	router.NotFoundHandler = xrouter.DevelopHandler()
}
