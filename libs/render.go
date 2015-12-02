package libs

import "net/http"

// Render struct
type Render struct{}

// Render function handler
func (r *Render) Render(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
