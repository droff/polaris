package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// HomeIndex GET "/"
func HomeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	t := template.Must(template.ParseFiles("views/layouts/application.htm", "views/login/login.htm"))

	if err := t.Execute(w, &Access{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PageNotFound GET "/"
func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	responseText := fmt.Sprintf("'%s' not found", mux.Vars(r)["url"])

	json.NewEncoder(w).Encode(Params{"status": 404, "data": responseText})
}
