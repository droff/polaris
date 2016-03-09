package lib

import (
	"log"
	"net/http"
	"time"
)

// Params contains json data
var Params map[string]interface{}

// Dispatcher func
func Dispatcher(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.RemoteAddr = r.Header.Get("X-Forwarded-For")

		w.Header().Set("Server", "polaris")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, time.Since(start))
	})
}
