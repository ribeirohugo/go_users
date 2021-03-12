package test

import (
	"github.com/ribeirohugo/go_users/internal/model"
)

const name = "name"
const password = "password"
const email = "email@domain"
const phone = "123456789"
const timestamp = 20

const nameInvalid = "na"
const passwordInvalid = "pass"
const emailInvalid = "mail"
const phoneInvalid = "phone"

const emailRepeated = "email2@domain"
const phoneRepeated = "012345678"

var user1 = model.User{Name: name, Password: password, Email: emailRepeated, Phone: phoneRepeated, Timestamp: timestamp}
var user2 = model.User{Name: name, Password: password, Email: "email3@domain", Phone: "001234567", Timestamp: timestamp}
var users = []model.User{user1, user2}

var user3 = model.User{Name: name, Password: password, Email: "email4@domain", Phone: "000123456", Timestamp: timestamp}

var backupUser1 = model.User{Name: name, Password: password, Email: emailRepeated, Phone: phoneRepeated, Timestamp: timestamp}
var backupUser2 = model.User{Name: name, Password: password, Email: "email3@domain", Phone: "001234567", Timestamp: timestamp}

func Init() {
	user1 = backupUser1
	user2 = backupUser2
	users = []model.User{user1, user2}
}
