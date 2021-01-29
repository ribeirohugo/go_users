package test

import (
	"../model"
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

var user = model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

func TestUserName(t *testing.T) {
	if user.Name != name {
		t.Errorf("User name incorrect, got: %s, want: %s.", user.Name, name)
	}
}

func TestPassword(t *testing.T) {
	if user.Password != password {
		t.Errorf("User password incorrect, got: %s, want: %s.", user.Password, password)
	}
}

func TestEmail(t *testing.T) {
	if user.Email != email {
		t.Errorf("User email incorrect, got: %s, want: %s.", user.Email, email)
	}
}

func TestPhone(t *testing.T) {
	if user.Phone != phone {
		t.Errorf("User phone incorrect, got: %s, want: %s.", user.Phone, phone)
	}
}

func TestTimestamp(t *testing.T) {
	if user.Timestamp != timestamp {
		t.Errorf("User timestamp incorrect, got: %d, want: %d.", user.Timestamp, timestamp)
	}
}

func TestInvalidName(t *testing.T) {
	u := model.User{Name: nameInvalid, Password: password, Email: email, Phone: phone, Timestamp: timestamp}
	flag := u.IsValid()
	if flag {
		t.Errorf("User name should return invalid. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidPassword(t *testing.T) {
	u := model.User{Name: name, Password: passwordInvalid, Email: email, Phone: phone, Timestamp: timestamp}
	flag := u.IsValid()
	if flag {
		t.Errorf("User password should return invalid. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidEmail(t *testing.T) {
	u := model.User{Name: name, Password: password, Email: emailInvalid, Phone: phone, Timestamp: timestamp}
	flag := u.IsValid()
	if flag {
		t.Errorf("User email should return invalid. Got: %t, want: %t.", flag, false)
	}
}

func TestInvalidPhone(t *testing.T) {
	u := model.User{Name: name, Password: password, Email: email, Phone: phoneInvalid, Timestamp: timestamp}
	flag := u.IsValid()
	if flag {
		t.Errorf("User phone should be invalid. Got: %t, want: %t.", flag, false)
	}
}
