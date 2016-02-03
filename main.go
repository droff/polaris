package main

import (
	"log"
	"net/http"

	"github.com/droff/polaris/lib"
	"github.com/droff/polaris/models"
)

func init() {
	models.InitDB()
}

func main() {
	router := lib.NewRouter()

	log.Println("Listening...")
	http.ListenAndServe(":4000", router)
}
