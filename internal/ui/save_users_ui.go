package ui

import (
	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/model"
)

func SaveUsersUI(users []model.User) {
	controller.SaveUsersController(users)
}
