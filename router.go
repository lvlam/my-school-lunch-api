package main

import (
	"net/http"

	gmux "github.com/gorilla/mux"
)

func NewRouter() *gmux.Router {

	router := gmux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
