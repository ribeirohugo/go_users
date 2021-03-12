package controller

import (
	"github.com/ribeirohugo/go_users/internal/model"
)

func FindUserController(email string, phone string, users *[]model.User) (*model.User, int) {

	if phone == "" && email == "" {
		return &model.User{}, -1
	} else if email == "" {
		return findUserByPhone(phone, users)
	} else if phone == "" {
		return findUserByEmail(email, users)
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
	return &model.User{}, -1
}

func findUserByEmail(email string, users *[]model.User) (*model.User, int) {
	var cont = 0

	for _, user := range *users {
		if user.Email == email {
			return &user, cont
		}
		cont++
	}
	return &model.User{}, -1
}

func findUserByPhone(phone string, users *[]model.User) (*model.User, int) {
	var cont = 0

	for _, user := range *users {
		if user.Phone == phone {
			return &user, cont
		}
		cont++
	}
	return &model.User{}, -1
}
