package ui

import (
	"../controller"
	"../model"
	"fmt"
)

func RemoveUserUI(users *[]model.User) {
	var email, phone string

	fmt.Println("Find user by:")

	fmt.Print("Email: (leave empty to ignore)")
	fmt.Scanf("%s", &email)

	fmt.Print("Phone: (leave empty to ignore)")
	fmt.Scanf("%s", &phone)

	_, position := controller.FindUserController(email, phone, users)

	if position < 0 {
		fmt.Println("User not found.")
	} else {
		controller.RemoveUserController(position, users)
		fmt.Println("User successfully removed.")
	}
}
