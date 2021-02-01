package test

import (
	"../../controller"
	"testing"
)

func TestFindUserByEmail(t *testing.T) {
	user, position := controller.FindUserController(emailRepeated, phone, &users)

	if *user != user1 {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *user, user1)
	}

	if position != 0 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, 0)
	}
}

func TestFindUserByPhone(t *testing.T) {
	user, position := controller.FindUserController(email, phoneRepeated, &users)

	if *user != user1 {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *user, user1)
	}

	if position != 0 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, 0)
	}
}

func TestFindUserByEmailAndPhone(t *testing.T) {
	user, position := controller.FindUserController(emailRepeated, phoneRepeated, &users)

	if *user != user1 {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *user, user1)
	}

	if position != 0 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, 0)
	}
}

func TestUserNotFound(t *testing.T) {
	_, position := controller.FindUserController(emailInvalid, phoneInvalid, &users)

	if position != -1 {
		t.Errorf("Invalid user not found position return. Got: %d, want: %d.", position, -1)
	}
}
