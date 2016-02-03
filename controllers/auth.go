package controllers

import (
	"encoding/json"
	"github.com/droff/polaris/models"
	"net/http"
)

// DataAuth struct
type DataAuth struct {
	Code string `json:"code"`
}

var codes models.Codes

// Session contains global session store
var Session models.Session

func init() {
	Session.Data = make(map[string]bool)
}

// Auth POST '/auth'
func Auth(w http.ResponseWriter, r *http.Request) {
	codes.Init()
	params := getParams(r)
	data := Params{}

	if codes.Find(params["code"].(string)) == true {
		data = Params{"status": "ok", "session_key": Session.Generate()}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

	json.NewEncoder(w).Encode(data)
}
