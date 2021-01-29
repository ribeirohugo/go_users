package controller

import (
	"../model"
)

func RemoveUserController(position int, users *[]model.User) {
	var array = *users

	array[position] = array[len(array)-1]
	*users = array[:len(array)-1]
}
