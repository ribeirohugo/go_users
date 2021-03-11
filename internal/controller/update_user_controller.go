package controller

import (
	"github.com/ribeirohugo/go_users/internal/model"
)

func UpdateUserController(name string, password string, email string, phone string, timestamp int, position int, users *[]model.User) bool {

	var usersArray = *users
	var user = model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

	if user.IsValid() {
		usersArray[position] = user
		users = &usersArray
		return true
	}

	return false
}
