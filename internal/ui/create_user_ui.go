package ui

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/model"
)

func CreateUserUI(users *[]model.User) {
	var name, password, email, phone string
	var timestamp int

	fmt.Print("Enter name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("Enter password: ")
	fmt.Scanf("%s", &password)

	fmt.Print("Enter email: ")
	fmt.Scanf("%s", &email)

	fmt.Print("Enter phone: ")
	fmt.Scanf("%s", &phone)

	fmt.Print("Enter timestamp: ")
	fmt.Scanf("%d", &timestamp)

	var flag bool
	var user model.User

	flag, user = controller.CreateUserController(name, password, email, phone, timestamp, users)

	if flag {
		fmt.Println("User: ", user)
		_, position := controller.FindUserController(email, phone, users)

		if position < 0 {
			fmt.Print("Email or phone already registered.")
		} else {
			fmt.Println("User successfully registered.")
		}
	} else {
		fmt.Println("Error registering user due to invalid data inserted.")
	}
}
