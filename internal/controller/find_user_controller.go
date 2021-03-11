package controller

import (
	"github.com/ribeirohugo/go_users/model"
)

func FindUserController(email string, phone string, users *[]model.User) (usr *model.User, position int) {

	if email == "" {
		return findUserByPhone(phone, users)
	} else if phone == "" {
		return findUserByEmail(email, users)
	} else if phone == "" && email == "" {
		return usr, -1
	}

	var cont = 0

	for _, user := range *users {
		if user.Email == email {
			return &user, cont
		} else if user.Phone == phone {
			return &user, cont
		}
		cont++
	}
	return usr, -1
}

func findUserByEmail(email string, users *[]model.User) (usr *model.User, position int) {
	var cont = 0

	for _, user := range *users {
		if user.Email == email {
			return &user, cont
		}
		cont++
	}
	return usr, -1
}

func findUserByPhone(phone string, users *[]model.User) (usr *model.User, position int) {
	var cont = 0

	for _, user := range *users {
		if user.Phone == phone {
			return &user, cont
		}
		cont++
	}
	return usr, -1
}
