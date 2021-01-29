package test

import (
	"../../controller"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	result := controller.UpdateUserController(name, password, email, phone, timestamp, 0, &users)

	if !result {
		t.Errorf("Invalid return user update. Got: %t, want: %t.", result, true)
	}
}

func TestUpdateUserInvalidName(t *testing.T) {
	result := controller.UpdateUserController(nameInvalid, password, email, phone, timestamp, 0, &users)

	if result {
		t.Errorf("Invalid return after invalid name inserted. Got: %t, want: %t.", result, false)
	}
}

func TestUpdateUserInvalidPassword(t *testing.T) {
	result := controller.UpdateUserController(nameInvalid, passwordInvalid, email, phone, timestamp, 0, &users)

	if result {
		t.Errorf("Invalid return after invalid password inserted. Got: %t, want: %t.", result, false)
	}
}

func TestUpdateUserInvalidEmail(t *testing.T) {
	result := controller.UpdateUserController(name, password, emailInvalid, phone, timestamp, 0, &users)

	if result {
		t.Errorf("Invalid return after invalid email inserted. Got: %t, want: %t.", result, false)
	}
}

func TestUpdateUserInvalidPhone(t *testing.T) {
	result := controller.UpdateUserController(name, password, email, phoneInvalid, timestamp, 0, &users)

	if result {
		t.Errorf("Invalid return after invalid phone inserted. Got: %t, want: %t.", result, false)
	}
}
