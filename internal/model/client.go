package model

import (
	"encoding/gob"
	"net"
	"sync"

	"github.com/ribeirohugo/go_users/internal/fault"
)

type Client struct {
	hostname string
	port     string
	message  string
	mutex    sync.Mutex
}

func (client *Client) SendUsers(users *[]User, network string, server string) bool {
	result := false

	client.mutex.Lock()

	con, err := net.Dial(network, server)
	fault.HandleFatalError("Error creating connection. ", err)

	if len(*users) > 0 {
		encoder := gob.NewEncoder(con)
		err := encoder.Encode(&users)
		fault.HandleError("Error encoding users slice. ", err)
		result = true
	}

	err = con.Close()
	fault.HandleFatalError("Error closing connection. ", err)

	client.mutex.Unlock()

	return result
}

func (client *Client) ReadUsers(users *[]User, con net.Conn) bool {
	result := false

	client.mutex.Lock()

	decoder := gob.NewDecoder(con)

	err := decoder.Decode(&users)
	fault.HandleError("Error decoding user. ", err)

	if len(*users) > 0 {
		result = true
	}

	client.mutex.Unlock()

	return result
}
