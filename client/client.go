package main

import (
	"../model"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server := getAddress(os.Args)

	con, err := net.Dial(network, server)
	handleError("Error creating connection. ", err)

	encoder := gob.NewEncoder(con)

	var name, password, email, phone, option string

	for {
		fmt.Println("Enter name: ")
		_, err = fmt.Scanf("%s", &name)
		handleError("Error entering name. ", err)

		fmt.Println("Enter password: ")
		_, err = fmt.Scanf("%s", &password)
		handleError("Error entering password. ", err)

		fmt.Println("Enter email: ")
		_, err = fmt.Scanf("%s", &email)
		handleError("Error entering email. ", err)

		fmt.Println("Enter phone: ")
		_, err = fmt.Scanf("%s", &phone)
		handleError("Error entering phone. ", err)

		timestamp := int(time.Now().Unix())

		user := model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

		if user.IsValid() {
			var users []model.User
			users = append(users, user)
			err = encoder.Encode(&users)
			handleError("Error encoding user. ", err)
		} else {
			fmt.Println("Invalid user.")
		}

		fmt.Println("Do you want to add another user? Press \"y\" to add.")
		_, err = fmt.Scanf("%s", &option)
		handleError("Error entering phone. ", err)

		if option != "y" {
			break
		}
	}

	err = con.Close()
	handleFatalError("Error closing connection. ", err)
}
