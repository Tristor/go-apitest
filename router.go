package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

//NewRouter gets called from main.go in order to simplify the dynamic
//generation of routes and provides a mux.Router that gets passed to
//http.ListenAndServe
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
