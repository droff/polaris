package main

import (
	"fmt"
	"net/http"

	"github.com/droff/polaris/controllers"
	"github.com/droff/polaris/libs"
	"github.com/droff/polaris/models"
)

// App configuration
const AppName string = "Polaris"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Polaris - %s", r.URL.Path[1:])
}

// Access controller
var Access controllers.Access

// Init method
func init() {
	var Users []models.User
	Users = append(Users, models.User{Username: "droff", Password: "12345", HasVideoAccess: true})
	Users = append(Users, models.User{Username: "test", Password: "test1", HasVideoAccess: false})

	Access.Title = "Login Page"
	Access.Users = &Users
}

func main() {
	fs := libs.Assets{http.Dir("assets/")}

	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(fs)))

	http.HandleFunc("/login", Access.Login)
	http.ListenAndServe(":4000", nil)
}
