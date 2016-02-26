package controllers

import (
	"encoding/base64"
	"github.com/droff/camsnap"
	"net/http"
	"strconv"
)

// GetImage sends bas64/image into stream
func GetImage(w http.ResponseWriter, r *http.Request) {
	params := getParams(r)
	if Session.Find(params["session_key"].(string)) {
		rawImage := camsnap.Get("/dev/video0", 1280, 720)
		imageEncoded := base64.StdEncoding.EncodeToString(rawImage)

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(imageEncoded)))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(imageEncoded))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
