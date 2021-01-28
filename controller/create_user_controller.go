package controller

import "../model"

func CreateUserController(name string, password string, email string, phone string, timestamp int) (bool, *model.User) {

	var user = model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

	if !user.IsValid() {
		return false, nil
	}

	return true, &user
}
