package ui

import (
	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/model"
)

func ReadUsersUI(users *[]model.User) {

	controller.ReadUsersController(users)

}
