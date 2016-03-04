package main

import (
	"log"
	"net/http"

	"github.com/droff/polaris/lib"
	"github.com/droff/polaris/models"
	"golang.org/x/net/http2"
)

func init() {
	models.InitDB()
}

func main() {
	router := lib.NewRouter()
	srvTLS := &http.Server{Addr: ":4000", Handler: router}

	go func() {
		srv := &http.Server{Addr: ":3000"}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://0.0.0.0:4000"+r.RequestURI, http.StatusMovedPermanently)
		})
		srv.ListenAndServe()
	}()

	log.Println("Listening...")
	http2.ConfigureServer(srvTLS, &http2.Server{})
	srvTLS.ListenAndServeTLS("openssl/server.crt", "openssl/server.key")
}
