package test

import (
	"testing"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/model"
)

const name = "name"
const password = "password"
const email = "email@domain"
const phone = "123456789"
const timestamp = 20

const nameInvalid = "na"
const passwordInvalid = "pass"
const emailInvalid = "mail"
const phoneInvalid = "phone"

const emailRepeated = "email2@domain"
const phoneRepeated = "012345678"

var user1 = model.User{Name: name, Password: password, Email: emailRepeated, Phone: phoneRepeated, Timestamp: timestamp}
var user2 = model.User{Name: name, Password: password, Email: "email3@domain", Phone: "001234567", Timestamp: timestamp}
var users = []model.User{user1, user2}

var user3 = model.User{Name: name, Password: password, Email: "email4@domain", Phone: "000123456", Timestamp: timestamp}

func TestInvalidName(t *testing.T) {
	flag, _ := controller.CreateUserController(nameInvalid, password, email, phone, timestamp, &users)

	if flag {
		t.Errorf("Invalid name wrong return. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidPassword(t *testing.T) {
	flag, _ := controller.CreateUserController(name, passwordInvalid, email, phone, timestamp, &users)

	if flag {
		t.Errorf("Invalid password wrong return. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidEmail(t *testing.T) {
	flag, _ := controller.CreateUserController(nameInvalid, password, emailInvalid, phone, timestamp, &users)

	if flag {
		t.Errorf("Invalid email wrong return. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidPhone(t *testing.T) {
	flag, _ := controller.CreateUserController(name, password, email, phoneInvalid, timestamp, &users)

	if flag {
		t.Errorf("Invalid phone wrong return. Got: %t, want: %t.", flag, false)
	}
}

func TestRepeatedUserEmail(t *testing.T) {
	flag, _ := controller.CreateUserController(name, password, emailRepeated, phone, timestamp, &users)

	if flag {
		t.Errorf("Repeated user email test. Got: %t, want: %t.", flag, false)
	}
}

func TestRepeatedUserPhone(t *testing.T) {
	flag, _ := controller.CreateUserController(name, password, email, phoneRepeated, timestamp, &users)

	if flag {
		t.Errorf("Repeated user email test. Got: %t, want: %t.", flag, false)
	}
}

func TestUserSuccessfullyCreated(t *testing.T) {
	expLength := len(users) + 1

	flag, user := controller.CreateUserController(name, password, email, phone, timestamp, &users)

	if !flag {
		t.Errorf("Error creating user. Got: %t, want: %t.", !flag, true)
	}

	if user.Name != name {
		t.Errorf("User name incorrect. Got: %s, want: %s.", user.Name, name)
	}

	if user.Password != password {
		t.Errorf("User password incorrect. Got: %s, want: %s.", user.Password, password)
	}

	if user.Email != email {
		t.Errorf("User email incorrect. Got: %s, want: %s.", user.Email, email)
	}

	if user.Phone != phone {
		t.Errorf("User phone incorrect. Got: %s, want: %s.", user.Phone, phone)
	}

	if user.Timestamp != timestamp {
		t.Errorf("User timestamp incorrect. Got: %d, want: %d.", user.Timestamp, timestamp)
	}

	currentLength := len(users)

	if expLength != currentLength {
		t.Errorf("Wrong expected length after creating user. Got: %d, want: %d.", currentLength, expLength)
	}
}

func TestAddUserFail(t *testing.T) {
	expLength := len(users)

	flag := controller.AddUserController(user1, &users)

	if flag {
		t.Errorf("Error failing to add user. Got: %t, want: %t.", !flag, false)
	}

	currentLength := len(users)

	if expLength != currentLength {
		t.Errorf("Wrong expected length failing to add user. Got: %d, want: %d.", currentLength, expLength)
	}
}

func TestAddUserSuccessfully(t *testing.T) {
	expLength := len(users) + 1

	flag := controller.AddUserController(user3, &users)

	if !flag {
		t.Errorf("Error failing to add user. Got: %t, want: %t.", !flag, true)
	}

	currentLength := len(users)

	if expLength != currentLength {
		t.Errorf("Wrong expected length after adding user. Got: %d, want: %d.", currentLength, expLength)
	}
}
