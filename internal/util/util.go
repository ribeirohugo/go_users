package util

import "log"

const (
	hostDefault = "localhost:"
	portDefault = "8080"
)

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

func GetAddress(arguments []string) string {
	var server string
	if len(arguments) >= 3 { //host + port
		server = arguments[1] + ":" + arguments[2]
	} else if len(arguments) == 2 { //port
		server = hostDefault + arguments[1]
	} else { //none
		server = hostDefault + portDefault
	}
	return server
}
