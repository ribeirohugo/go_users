package controller

import (
	"../model"
)

func CreateUserController(name string, password string, email string, phone string, timestamp int, users *[]model.User) (flag bool, usr model.User) {

	var user = model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

	if !user.IsValid() {
		return false, usr
	}

	FindUserController(email, phone, users)

	return true, user
}
