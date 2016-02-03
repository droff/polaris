package models

import (
	"github.com/droff/polaris/models"
	"testing"
)

var session models.Session

func init() {
}

func TestGenerate(t *testing.T) {
	h1 := session.Generate()
	h2 := session.Generate()

	if h1 == h2 {
		t.Error("h1 should not equal h2")
	}
}
