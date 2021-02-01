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

	var name, password, email, phone string

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

	fmt.Println("Do you want to add another user?")
	handleError("Error entering phone. ", err)

	timestamp := time.Now().Unix()

	user := model.User{Name: name, Password: password, Email: email, Phone: phone, Timestamp: int(timestamp)}

	if user.IsValid() {
		err = encoder.Encode(&user)
		handleError("Error encoding user. ", err)
	} else {
		fmt.Println("Invalid user.")
	}

	err = con.Close()
	handleError("Error closing connection. ", err)
}
