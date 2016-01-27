package controllers

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

// GetImage send image into stream
func GetImage(w http.ResponseWriter, r *http.Request) {
	takeSnap()

	rawImage, _ := ioutil.ReadFile("/tmp/shot.jpg")
	imageEncoded := base64.StdEncoding.EncodeToString(rawImage)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(imageEncoded)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(imageEncoded))
}

func takeSnap() {
	//cmd := exec.Command("scrot", "/tmp/shot.jpg")
	cmd := exec.Command("fswebcam", "-r", "1280x720", "--jpeg", "100", "--no-banner", "/tmp/shot.jpg")
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		log.Printf("fswebcam ERROR: %v\n", err)
	}
}
