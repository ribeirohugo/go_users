package test

import (
	"../model"
	"testing"
)

const name = "name"
const password = "password"
const email = "email"
const phone = "phone"
const timestamp = 20

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
