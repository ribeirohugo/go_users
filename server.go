package main

import (
	"./controller"
	"./model"
	"./util"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
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
		handleRequest(c, &database)
	}

	//err = con.Close()
	//handleFatalError("Error closing server. ", err)
}

func handleRequest(con net.Conn, database *[]model.User) {
	fmt.Printf("Serving %s\n", con.RemoteAddr().String())

	var users []model.User
	mutex := sync.Mutex{}
	mutex.Lock()

	data := gob.NewDecoder(con)

	err := data.Decode(&users)
	util.HandleError("Error decoding user. ", err)

	for i, _ := range users {
		user := users[i]

		if user.IsValid() {
			if controller.AddUserController(user, database) {
				fmt.Println(user, "User successfully saved.")
			} else {
				fmt.Println(user, "User already exists in database.")
			}
		} else {
			fmt.Println("Invalid user.")
		}
	}
	controller.SaveUsersController(*database)

	mutex.Unlock()
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
