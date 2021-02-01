package controller

import (
	"../model"
	"encoding/gob"
	"os"
)

func ReadUsersController(users *[]model.User) {
	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError("Error opening file.", err)

	dataDecoder := gob.NewDecoder(file)
	err = dataDecoder.Decode(&users)

	err = file.Close()
	checkError("Error closing file.", err)
}
