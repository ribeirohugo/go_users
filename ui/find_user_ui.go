package ui

import (
	"../controller"
	"../model"
	"fmt"
)

func FindUserUI(users *[]model.User) {

	var email, phone string

	fmt.Print("Enter email: ")
	fmt.Scanf("%s", &email)

	fmt.Print("Enter phone: ")
	fmt.Scanf("%s", &phone)

	user, result := controller.FindUserController(email, phone, users)

	if result {
		fmt.Println(user)
	} else {
		fmt.Println("No user found.")
	}
}
