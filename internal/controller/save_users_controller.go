package controller

import (
	"encoding/gob"
	"github.com/ribeirohugo/go_users/internal/fault"
	"os"

	"github.com/ribeirohugo/go_users/internal/model"
)

func SaveUsersController(users []model.User, binFile string) {
	file, err := os.Create(binFile)
	fault.HandleFatalError("Error creating file.", err)

	dataEncoder := gob.NewEncoder(file)
	dataEncoder.Encode(users)

	err = file.Close()
	fault.HandleFatalError("Error closing file.", err)
}
