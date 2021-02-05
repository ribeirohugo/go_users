package controller

import (
	"../model"
	"encoding/gob"
	"os"
)

func ReadUsersController(users *[]model.User) {
	file, err := os.Open(dataFile)
	checkError("Error opening file.", err)

	dataDecoder := gob.NewDecoder(file)
	err = dataDecoder.Decode(&users)

	err = file.Close()
	checkError("Error closing file.", err)
}
