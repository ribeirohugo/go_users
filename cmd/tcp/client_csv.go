package main

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleError("", err)

	con, err := net.Dial(network, cfg.HttpHost)
	fault.HandleFatalError("Error creating connection. ", err)

	file, err := os.Open(cfg.CsvPath)
	fault.HandleFatalError("Error opening Csv file.", err)
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	fault.HandleError("Error reading Csv lines.", err)

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
		fault.HandleError("Error encoding users slice. ", err)
	}

	err = con.Close()
	fault.HandleFatalError("Error closing connection. ", err)
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
