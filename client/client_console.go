package main

import (
	"encoding/gob"
	"fmt"
	util2 "github.com/ribeirohugo/go_users/internal/fault"
	model2 "github.com/ribeirohugo/go_users/internal/model"
	"net"
	"os"
	"time"
)

func main() {
	server := getAddress(os.Args)

	con, err := net.Dial(network, server)
	util2.HandleError("Error creating connection. ", err)

	encoder := gob.NewEncoder(con)

	var name, password, email, phone, option string
	var users []model2.User

	for {
		fmt.Println("Enter name: ")
		_, err = fmt.Scanf("%s", &name)
		util2.HandleError("Error entering name. ", err)

		fmt.Println("Enter password: ")
		_, err = fmt.Scanf("%s", &password)
		util2.HandleError("Error entering password. ", err)

		fmt.Println("Enter email: ")
		_, err = fmt.Scanf("%s", &email)
		util2.HandleError("Error entering email. ", err)

		fmt.Println("Enter phone: ")
		_, err = fmt.Scanf("%s", &phone)
		util2.HandleError("Error entering phone. ", err)

		timestamp := int(time.Now().Unix())

		user := model2.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: timestamp}

		if user.IsValid() {
			users = append(users, user)
		} else {
			fmt.Println("Invalid user.")
		}

		fmt.Println("Do you want to add another user? Press \"y\" to add.")
		_, err = fmt.Scanf("%s", &option)
		util2.HandleError("Error entering phone. ", err)

		if option != "y" {
			break
		}
	}

	noUsers := len(users)
	if noUsers > 0 {
		err = encoder.Encode(&users)
		util2.HandleError("Error encoding user. ", err)
		fmt.Println(noUsers, " users sent.")
	} else {
		fmt.Println("No user was sent.")
	}

	err = con.Close()
	util2.HandleFatalError("Error closing connection. ", err)

	fmt.Println("Client closed.")
}
