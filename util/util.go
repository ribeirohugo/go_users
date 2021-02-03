package util

import "log"

func HandleFatalError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}

func HandleError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
