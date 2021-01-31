package main

import (
	"./model"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	network     = "tcp"
	hostDefault = "localhost:"
	portDefault = "8080"
)

func main() {
	server := getAddress(os.Args)

	con, err := net.Listen(network, server)
	handleError("Error creating server. ", err)

	for {
		c, err := con.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		handleRequest(c)
	}

	err = con.Close()
	handleError("Error closing server. ", err)
}

func handleRequest(con net.Conn) {
	fmt.Printf("Serving %s\n", con.RemoteAddr().String())

	data := gob.NewDecoder(con)
	user := &model.User{}
	err := data.Decode(user)
	handleError("Error decoding user. ", err)

	fmt.Println()
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

func handleError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
