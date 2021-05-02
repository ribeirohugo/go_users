package main

import (
	"fmt"

	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
	"github.com/ribeirohugo/go_users/internal/ui"
)

const (
	configFile     = "config.toml"
	options    int = 8
)

var users []model.User

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleError("", err)

	var option = 0

	for option != options {
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
			ui.ReadUsersUI(&users, cfg.BinPath)
		case 7:
			ui.SaveUsersUI(users, cfg.BinPath)
		case options:
			break
		default:
			fmt.Println("Invalid option.")
		}
	}

	fmt.Println("Program exited.")
}
