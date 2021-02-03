package main

import (
	"./controller"
	"./model"
	"./util"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

const (
	network     = "tcp"
	hostDefault = "localhost:"
	portDefault = "8080"
)

var database []model.User

func main() {
	fmt.Println("Server started.")

	controller.ReadUsersController(&database)

	server := getAddress(os.Args)

	con, err := net.Listen(network, server)
	util.HandleFatalError("Error creating server. ", err)
	defer con.Close()
	util.HandleFatalError("Error closing server. ", err)

	for {
		c, err := con.Accept()
		util.HandleError("Error accepting new connection.", err)
		handleRequest(c)
	}

	//err = con.Close()
	//handleFatalError("Error closing server. ", err)
}

func handleRequest(con net.Conn) {
	fmt.Printf("Serving %s\n", con.RemoteAddr().String())

	data := gob.NewDecoder(con)

	var users []model.User

	err := data.Decode(&users)
	util.HandleError("Error decoding user. ", err)

	for i, _ := range users {
		user := users[i]

		if user.IsValid() {
			if controller.AddUserController(user, &database) {
				fmt.Println(user, "User successfully saved.")
			} else {
				fmt.Println(user, "User already exists in database.")
			}
		} else {
			fmt.Println("Invalid user.")
		}
	}
	controller.SaveUsersController(database)
}

func getAddress(arguments []string) string {
	var server string
	if len(arguments) >= 3 { //host + port
		server = arguments[1] + ":" + arguments[2]
	} else if len(arguments) == 2 { //port
		server = hostDefault + arguments[1]
	} else { //none
		server = hostDefault + portDefault
	}
	return server
}
