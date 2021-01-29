package test

import (
	"../../controller"
	"testing"
)

func TestRemoveUser(t *testing.T) {
	initialLength := len(users)

	result := controller.RemoveUserController(0, &users)

	difference := initialLength - len(users)
	expDifference := 1

	if difference != expDifference {
		t.Errorf("Invalid length after removing user. Got: %d, want: %d.", difference, expDifference)
	}

	if !result {
		t.Errorf("Invalid remove user return. Got: %t, want: %t.", result, true)
	}
}

func TestRemoveUserInvalidPosition(t *testing.T) {
	initialLength := len(users)

	result := controller.RemoveUserController(len(users)+1, &users)

	finalLength := len(users)

	if finalLength != initialLength {
		t.Errorf("Invalid length after removing user. Got: %d, want: %d.", finalLength, initialLength)
	}

	if result {
		t.Errorf("Invalid remove user return. Got: %t, want: %t.", result, false)
	}
}
