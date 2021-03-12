package test

import (
	"testing"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/model"
)

const binFile = "../users.bin"

func TestSaveReadUsers(t *testing.T) {
	Init()

	expected := 2
	length := len(users)

	if length != expected {
		t.Errorf("Invalid read users length. Got: %d, want: %d.", length, expected)
	}

	controller.SaveUsersController(users, binFile)

}

func TestReadUsers(t *testing.T) {
	var readUsers []model.User

	controller.ReadUsersController(&readUsers, binFile)

	expected := 2
	length := len(readUsers)

	if length != expected {
		t.Errorf("Invalid read users length. Got: %d, want: %d.", length, expected)
	}
}
