package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ribeirohugo/go_users/internal/controller"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func apiReader(w http.ResponseWriter, req *http.Request) {

	var inputUsers []model.User
	var database []model.User
	controller.ReadUsersController(&database)

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

	controller.SaveUsersController(database)

	fmt.Println("Body: ", bodyStr)
	fmt.Println("-------------------")
}

func main() {
	address := fault.GetAddress(os.Args)

	http.HandleFunc("/api", apiReader)
	err := http.ListenAndServe(address, nil)

	fault.HandleFatalError("", err)
}
