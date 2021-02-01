package controller

import (
	"../model"
)

func CreateUserController(name string, password string, email string, phone string, timestamp int, users *[]model.User) (flag bool, usr model.User) {

	var user = model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

	if !user.IsValid() {
		return false, usr
	}

	_, position := FindUserController(email, phone, users)
	if position < 0 {
		return false, usr
	}

	*users = append(*users, user)
	return true, user
}

func AddUserController(user model.User, users *[]model.User) bool {
	if !user.IsValid() {
		return false
	}

	_, position := FindUserController(user.Email, user.Phone, users)
	if position >= 0 {
		return false
	}

	*users = append(*users, user)
	return true
}
