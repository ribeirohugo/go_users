package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
	"net/http"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleError("", err)

	http.HandleFunc("/api", apiReader)
	err = http.ListenAndServe(cfg.HttpHost, nil)

	fault.HandleFatalError("", err)
}

func apiReader(w http.ResponseWriter, req *http.Request) {
	cfg, err := config.Load(configFile)
	fault.HandleFatalError("", err)

	var inputUsers []model.User
	var database []model.User
	controller.ReadUsersController(&database, cfg.BinPath)

	//Write in console
	fmt.Println("Remote Address: ", req.RemoteAddr)
	fmt.Println("User-Agent: ", req.UserAgent())
	fmt.Println("Method: ", req.Method)

	buf := new(bytes.Buffer)

	buf.ReadFrom(req.Body)
	bodyStr := buf.String()

	json.Unmarshal(buf.Bytes(), &inputUsers)
	for _, user := range inputUsers {
		controller.AddUserController(user, &database)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[{\"Status\": \"Ok\"}]"))

	controller.SaveUsersController(database, cfg.BinPath)

	fmt.Println("Body: ", bodyStr)
	fmt.Println("-------------------")
}
