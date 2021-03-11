package controller

import (
	"github.com/ribeirohugo/go_users/internal/model"
)

func RemoveUserController(position int, users *[]model.User) bool {
	var array = *users

	if position >= 0 && position < len(array) {
		array[position] = array[len(array)-1]
		*users = array[:len(array)-1]
		return true
	}

	return false
}
