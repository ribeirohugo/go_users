package test

import (
	"../../controller"
	"testing"
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

func TestUserSucessfullyCreated(t *testing.T) {
	flag, user := controller.CreateUserController(name, password, email, phone, timestamp)

	if !flag {
		t.Errorf("Error creating user. Got: %t, want: %t.", !flag, flag)
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
}

func TestInvalidName(t *testing.T) {
	flag, _ := controller.CreateUserController(nameInvalid, password, email, phone, timestamp)

	if flag {
		t.Errorf("Invalid name wrong return. Got: %t, want: %t.", flag, !flag)
	}
}

func TestInvalidPassword(t *testing.T) {
	flag, _ := controller.CreateUserController(name, passwordInvalid, email, phone, timestamp)

	if flag {
		t.Errorf("Invalid password wrong return. Got: %t, want: %t.", flag, !flag)
	}
}

func TestInvalidEmail(t *testing.T) {
	flag, _ := controller.CreateUserController(nameInvalid, password, emailInvalid, phone, timestamp)

	if flag {
		t.Errorf("Invalid email wrong return. Got: %t, want: %t.", flag, !flag)
	}
}

func TestInvalidPhone(t *testing.T) {
	flag, _ := controller.CreateUserController(name, password, email, phoneInvalid, timestamp)

	if flag {
		t.Errorf("Invalid phone wrong return. Got: %t, want: %t.", flag, !flag)
	}
}
