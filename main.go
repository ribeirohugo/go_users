package main

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/ui"
	"github.com/ribeirohugo/go_users/model"
)

const OPTIONS int = 8

var users []model.User

func main() {

	var option = 0

	for option != OPTIONS {

		fmt.Println("Select an option:\n" +
			"1 - Create user\n" +
			"2 - List users\n" +
			"3 - Find user\n" +
			"4 - Update user\n" +
			"5 - Remove user\n" +
			"6 - Load data\n" +
			"7 - Save data\n" +
			"8 - Exit")

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			ui.CreateUserUI(&users)
		case 2:
			ui.PrintUsersUI(users)
		case 3:
			ui.FindUserUI(&users)
		case 4:
			ui.UpdateUserUI(&users)
		case 5:
			ui.RemoveUserUI(&users)
		case 6:
			ui.ReadUsersUI(&users)
		case 7:
			ui.SaveUsersUI(users)
		case OPTIONS:
			break
		default:
			fmt.Println("Invalid option.")
		}
	}

	fmt.Println("Program exited.")
}
