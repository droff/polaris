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

	log.Println("Listening...")
	/*http.ListenAndServe(":4000", router)*/

	srv := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	http2.ConfigureServer(srv, &http2.Server{})
	srv.ListenAndServeTLS("openssl/server.crt", "openssl/server.key")
}
