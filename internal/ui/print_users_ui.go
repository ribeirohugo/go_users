package ui

import (
	"fmt"

	"github.com/ribeirohugo/go_users/model"
)

func PrintUsersUI(user []model.User) {
	for i := 0; i < len(user); i++ {
		fmt.Println(user[i])
	}
}
