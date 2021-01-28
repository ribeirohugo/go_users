package ui

import (
	"../controller"
	"../model"
	"fmt"
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
	var user *model.User

	flag, user = controller.CreateUserController(name, password, email, phone, timestamp)

	if flag {
		var userObj = *user
		fmt.Println("User: ", userObj)
		*users = append(*users, userObj)
	} else {
		fmt.Println("Erro ao registar utilizador.")
	}
}
