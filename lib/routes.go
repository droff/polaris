package lib

import (
	ctrls "github.com/droff/polaris/controllers"
	"net/http"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes list
type Routes []Route

var routes = Routes{
	Route{"HomeIndex", "GET", "/", ctrls.HomeIndex},
	Route{"Auth", "POST", "/auth", ctrls.Auth},
	Route{"GetImage", "POST", "/getimage", ctrls.GetImage},

	Route{"Error404", "GET", "/{url}", ctrls.PageNotFound},
}
