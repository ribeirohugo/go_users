package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"sync"

	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

const network = "tcp"

var database []model.User

func main() {
	cfg, err := config.Load()
	fault.HandleFatalError("", err)

	fmt.Println("Server started.")

	controller.ReadUsersController(&database, cfg.BinPath)

	con, err := net.Listen(network, cfg.TcpHost)
	fault.HandleFatalError("Error creating server. ", err)
	defer con.Close()

	for {
		c, err := con.Accept()
		fault.HandleError("Error accepting new connection.", err)
		handleRequest(c, &database, cfg.BinPath)
	}

	//err = con.Close()
	//handleFatalError("Error closing server. ", err)
}

func handleRequest(con net.Conn, database *[]model.User, binFile string) {
	fmt.Printf("Serving %s\n", con.RemoteAddr().String())

	var users []model.User
	mutex := sync.Mutex{}
	mutex.Lock()

	data := gob.NewDecoder(con)

	err := data.Decode(&users)
	fault.HandleError("Error decoding user. ", err)

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
	controller.SaveUsersController(*database, binFile)

	mutex.Unlock()
}
