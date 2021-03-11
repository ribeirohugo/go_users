package controller

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/ribeirohugo/go_users/internal/model"
)

func SaveUsersController(users []model.User, binFile string) {
	file, err := os.Create(binFile)
	checkError("Error creating file.", err)

	dataEncoder := gob.NewEncoder(file)
	dataEncoder.Encode(users)

	err = file.Close()
	checkError("Error closing file.", err)
}

func checkError(message string, err error) bool {
	if err != nil {
		log.Fatal(message, err)
		return true
	}
	return false
}
