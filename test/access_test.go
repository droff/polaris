package controllers

import (
	"testing"

	ctrl "github.com/droff/polaris/controllers"
	mdl "github.com/droff/polaris/models"
)

var Users []mdl.User
var Access ctrl.Access

func init() {
	Users = append(Users, mdl.User{Username: "test", Password: "test123"})
	Access.Users = &Users
}

func TestAuth(t *testing.T) {
	result := Access.Auth("test", "test123")
	if result == false {
		t.Error("Expected true, got ", result)
	}

	result = Access.Auth("test", "test1")
	if result == true {
		t.Error("Expected true, got ", result)
	}
}
