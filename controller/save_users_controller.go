package controller

import (
	"../model"
	"encoding/gob"
	"log"
	"os"
)

const dataFile string = "users.bin"

func SaveUsersController(users []model.User) {

	file, err := os.Create(dataFile)
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
