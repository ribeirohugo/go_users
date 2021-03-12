package test

import (
	"github.com/ribeirohugo/go_users/internal/model"
	"testing"

	"github.com/ribeirohugo/go_users/internal/controller"
)

func TestFindUserByEmail(t *testing.T) {
	user, position := controller.FindUserController(emailRepeated, "", &users)

	if *user != user1 {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *user, user1)
	}

	if position != 0 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, 0)
	}
}

func TestFindUserByEmailFail(t *testing.T) {
	Init()

	_, position := controller.FindUserController(email, "", &users)

	if position != -1 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, -1)
	}
}

func TestFindUserByPhone(t *testing.T) {
	user, position := controller.FindUserController("", phoneRepeated, &users)

	if *user != user1 {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *user, user1)
	}

	if position != 0 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, 0)
	}
}

func TestFindUserByPhoneFail(t *testing.T) {
	Init()

	_, position := controller.FindUserController("", phone, &users)

	if position != -1 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, -1)
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

func TestFindUserByEmailAndPhoneFail(t *testing.T) {
	Init()

	usr, position := controller.FindUserController("", "", &users)

	if position != -1 {
		t.Errorf("Invalid position return. Got: %d, want: %d.", position, -1)
	}

	expected := model.User{}
	if *usr != expected {
		t.Errorf("Invalid user return. Got: %q, want: %q.", *usr, expected)
	}
}

func TestUserNotFound(t *testing.T) {
	_, position := controller.FindUserController(emailInvalid, phoneInvalid, &users)

	if position != -1 {
		t.Errorf("Invalid user not found position return. Got: %d, want: %d.", position, -1)
	}

	_, position = controller.FindUserController("", "", &users)

	if position != -1 {
		t.Errorf("Invalid user not found position return. Got: %d, want: %d.", position, -1)
	}
}
