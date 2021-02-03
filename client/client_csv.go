package main

import (
	"../model"
	"../util"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args

	server := getAddress(arguments)

	con, err := net.Dial(network, server)
	util.HandleFatalError("Error creating connection. ", err)

	if len(arguments) < 3 {
		fmt.Println("Invalid file path.")
		return
	}

	path := arguments[3]

	file, err := os.Open(path)
	util.HandleFatalError("Error opening file.", err)
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	util.HandleError("Error reading csv lines.", err)

	var users []model.User

	for i, _ := range lines {
		flag, user := handleLine(lines[i][0], &users)

		if flag {
			fmt.Println(user)
		} else {
			fmt.Println("Invalid user.")
		}
	}

	if len(users) > 0 {
		encoder := gob.NewEncoder(con)
		err := encoder.Encode(&users)
		util.HandleError("Error encoding users slice. ", err)
	}

	err = con.Close()
	util.HandleFatalError("Error closing connection. ", err)
}

func handleLine(line string, users *[]model.User) (flag bool, usr model.User) {
	fields := strings.Split(line, ";")

	if len(fields) > 3 {
		timestamp := int(time.Now().Unix())
		user := model.User{Name: fields[0], Password: fields[1], Email: fields[2], Phone: fields[3], Timestamp: timestamp}

		if user.IsValid() {
			*users = append(*users, user)
			return true, user
		}
	}

	return false, usr
}
