package ui

import (
	"../controller"
	"../model"
)

func SaveUsersUI(users []model.User) {

	controller.SaveUsersController(users)

}
