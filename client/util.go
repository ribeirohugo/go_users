package main

const (
	network     = "tcp"
	hostDefault = "localhost:"
	portDefault = "8080"
)

func getAddress(arguments []string) string {
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
