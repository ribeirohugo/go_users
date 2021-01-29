package ui

import (
	"../model"
	"fmt"
)

func PrintUsersUI(user []model.User) {
	for i := 0; i < len(user); i++ {
		fmt.Println(user[i])
	}
}
