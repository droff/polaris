package controllers

import (
	"html/template"
	"net/http"

	"github.com/droff/polaris/models"
)

// Access params
type Access struct {
	Title     string
	LoginPage *LoginPage
	Users     *models.Users
}

// LoginPage page
type LoginPage struct {
	UserName       string
	HasVideoAccess bool
}

// Auth users
func (ctrl *Access) Auth(username, password string) *models.User {
	for _, user := range *ctrl.Users {
		if (user.Username == username) && (user.Password == password) {
			return &user
		}
	}

	return nil
}

// Login template
func (ctrl *Access) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("views/layouts/application.htm", "views/login/login.htm"))

		if err := t.Execute(w, &Access{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		r.ParseForm()

		user := ctrl.Auth(r.Form["username"][0], r.Form["password"][0])
		if user != nil {
			t := template.Must(template.ParseFiles("views/layouts/application.htm", "views/login/info.htm"))

			page := &Access{LoginPage: &LoginPage{UserName: r.Form["username"][0], HasVideoAccess: user.HasVideoAccess}}
			if err := t.Execute(w, page); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			http.Error(w, "Auth", http.StatusForbidden)
		}
	}
}
