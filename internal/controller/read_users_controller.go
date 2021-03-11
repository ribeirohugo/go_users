package controller

import (
	"encoding/gob"
	"os"

	"github.com/ribeirohugo/go_users/internal/model"
)

func ReadUsersController(users *[]model.User, binFile string) {
	file, err := os.Open(binFile)
	checkError("Error opening file.", err)

	dataDecoder := gob.NewDecoder(file)
	err = dataDecoder.Decode(&users)

	err = file.Close()
	checkError("Error closing file.", err)
}
