package fault

import "log"

const (
	hostDefault = "localhost:"
	portDefault = "8080"
)

func HandleFatalError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func HandleError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}
