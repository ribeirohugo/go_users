package model

type User struct {
	Name      string
	Password  string
	Email     string
	Phone     string
	Timestamp int
}

func (u User) IsValid() bool {
	var flag = true

	if len(u.Name) < 2 || len(u.Name) > 255 {

		flag = false
	}

	if u.Email == "" {

		flag = false
	}

	if len(u.Password) < 6 || len(u.Password) > 24 {

		flag = false
	}

	return flag
}
