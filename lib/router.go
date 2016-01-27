package lib

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter handler
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Dispatcher(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
