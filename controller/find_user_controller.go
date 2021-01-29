package controller

import "../model"

func FindUserController(email string, phone string, users *[]model.User) (usr model.User, result bool) {

	if email == "" {
		return findUserByPhone(phone, users)
	} else if phone == "" {
		return findUserByEmail(email, users)
	}

	for _, user := range *users {
		if user.Email == email {
			return user, true
		} else if user.Phone == phone {
			return user, true
		}
	}
	return usr, false
}

func findUserByEmail(email string, users *[]model.User) (usr model.User, result bool) {
	for _, user := range *users {
		if user.Email == email {
			return user, true
		}
	}
	return usr, false
}

func findUserByPhone(phone string, users *[]model.User) (usr model.User, result bool) {
	for _, user := range *users {
		if user.Phone == phone {
			return user, true
		}
	}
	return usr, false
}
