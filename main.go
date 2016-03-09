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

	log.Println("Listening...")
	http2.ConfigureServer(srvTLS, &http2.Server{})
	srvTLS.ListenAndServeTLS("openssl/server.crt", "openssl/server.key")
}
