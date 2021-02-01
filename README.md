# Go - Simple Users
**Simple Users** is a very simple user management tool made using Go.
It allows to create, update, remove and list users, persist them to a binary file.
You can use a console to manage data, or a server that receives information through a TCP connection.

## Run server

You can run your server using the following command and replacing `hostname` and `port` values:

``run go server.go hostname port``

You can also run the server by just inserting `port` value assuming hostname will be `localhost`:

`run go server.go port`

## Run client

Client can be used independently of main project, in order to send information to server through TCP.
There are two built-in examples using go in client folder: the console mode and the CSV reader mode. To run console mode run the following command:

`run go client/client.go client/util.go hostname port`

There is also a CSV loader to send users through TCP executing the following command:

``go run client/client_csv.go client/util.go hostname port file.csv``
