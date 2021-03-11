package ui

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/model"
)

func FindUserUI(users *[]model.User) {

	var email, phone string

	fmt.Print("Enter email: ")
	fmt.Scanf("%s", &email)

	fmt.Print("Enter phone: ")
	fmt.Scanf("%s", &phone)

	user, position := controller.FindUserController(email, phone, users)

	if position >= 0 {
		fmt.Println(*user)
	} else {
		fmt.Println("No user found.")
	}
}
