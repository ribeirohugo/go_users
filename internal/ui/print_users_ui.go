package ui

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/model"
)

func PrintUsersUI(user []model.User) {
	for i := 0; i < len(user); i++ {
		fmt.Println(user[i])
	}
}
