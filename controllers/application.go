package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Params builds new map
type Params map[string]interface{}

func getParams(r *http.Request) map[string]interface{} {
	params := Params{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &params)

	return params
}
