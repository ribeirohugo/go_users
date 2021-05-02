package controller

import (
	"encoding/gob"
	"os"

	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func ReadUsersController(users *[]model.User, binFile string) {
	file, err := os.Open(binFile)
	fault.HandleFatalError("Error opening file.", err)

	dataDecoder := gob.NewDecoder(file)
	err = dataDecoder.Decode(&users)

	err = file.Close()
	fault.HandleFatalError("Error closing file.", err)
}
