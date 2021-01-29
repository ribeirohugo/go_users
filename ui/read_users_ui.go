package ui

import (
	"../controller"
	"../model"
)

func ReadUsersUI(users *[]model.User) {

	controller.ReadUsersController(users)

}
