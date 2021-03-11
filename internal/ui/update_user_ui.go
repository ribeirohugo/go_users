package ui

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/model"
)

func UpdateUserUI(users *[]model.User) {
	var email, phone string

	fmt.Println("Find user by:")

	fmt.Print("Email: (leave empty to ignore)")
	fmt.Scanf("%s", &email)

	fmt.Print("Phone: (leave empty to ignore)")
	fmt.Scanf("%s", &phone)

	var user, position = controller.FindUserController(email, phone, users)

	if position < 0 {
		fmt.Println("User not found.")
	} else {

		fmt.Println("User: ", *user)
		name := user.Name
		password := user.Password
		email := user.Email
		phone := user.Phone
		timestamp := user.Timestamp

		fmt.Print("Enter name:  (leave empty to ignore)")
		fmt.Scanf("%s", &name)

		fmt.Print("Enter password:  (leave empty to ignore)")
		fmt.Scanf("%s", &password)

		fmt.Print("Enter email:  (leave empty to ignore)")
		fmt.Scanf("%s", &email)

		fmt.Print("Enter phone:  (leave empty to ignore)")
		fmt.Scanf("%s", &phone)

		fmt.Print("Enter timestamp:  (leave empty to ignore)")
		fmt.Scanf("%d", &timestamp)

		_, position := controller.FindUserController(email, phone, users)

		if position < 0 {
			result := controller.UpdateUserController(name, password, email, phone, timestamp, position, users)

			if result {
				fmt.Println("User successfully updated.")
			} else {
				fmt.Println("Invalid user data inserted.")
			}
		} else {
			fmt.Println("User email or phone already registered.")
		}
	}
}
