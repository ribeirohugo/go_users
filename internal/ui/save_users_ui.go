package ui

import (
	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/model"
)

func SaveUsersUI(users []model.User) {
	controller.SaveUsersController(users)
}
