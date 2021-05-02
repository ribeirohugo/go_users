# Go - Simple Users
**Simple Users** is a very simple user management tool made using Go.
It allows to create, update, remove and list users, persist them to a binary file.
You can use a console to manage data, or a server that receives information through a TCP connection.

This project is the first one I'm creating using Go and it aims to be a way to explore Go functionalities and learn more about this language.
It will be improved as long as I will get more expertise programming with Go.

## 1. Configurations

Fulfill the configuration file ``config.toml`` with the needed setups in order to use Simple Users different applications.

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``bin_path`` | Main binary file to persist user data. | `string` | `users.bin` | **NO** |
| ``csv_path`` | Csv file path to be used in tcp client. | `string` | `users.csv` | **NO** |
| ``http_host`` | Http host address to be used in http server app. | `string` | ` ` | **YES** |
| ``tcp_host`` | Tcp host address to be used in tcp server. | `string` | ` ` | **YES** |
| ``[mysql]`` | MySql database config data. | `Db` | ` ` | **YES** |
| ``[postgres]`` | Postgres database config data. | `Db` | ` ` | **YES** |

### 1.1. Database type

To set up ``[mysql]`` and ``[postgres]`` use the following parameters:

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``db`` | Database name. | `string` | ` ` | **YES** |
| ``host`` | Database host. | `string` | ` ` | **YES** |
| ``password`` | Database password. | `string` | ` ` | **YES** |
| ``port`` | Database port. | `int` | ` ` | **YES** |
| ``user`` | Database user with needed privileges over database. | `string` | ` ` | **YES** |

## 2. Simple Console

Manage users by using a menu console with all the options you need: since load a file, to execute CRUD operations.
To use Simple Console just run ``make run`` or the following code:

```
run go ./cmd/tcp/server.go
```

## 3. Tcp Application

### 3.1. Run server

To run Tcp Application you may firstly run Go Users server by run ``make tcp-server`` or the following command:

```
run go ./cmd/tcp/server.go
```

### 3.2. Tcp Client Console

Using client console allows inserting new users into the binary through command line.

To run console mode run ``make tcp-client-console`` or the following command:

```
run go ./cmd/tcp/client_console.go
```

### 3.3. Tcp Csv Client

There is also a CSV loader to send users through a TCP connection to server.
Fill the config file ``csv_path`` attribute Csv to the Csv file you want to send.
Csv rows should follow this pattern:

``<Name>;<Password>;<Email>;<Phone>``

To run Csv Client run make ``tcp-client-csv`` or the following command:

```
go run ./cmd/tcp/client_csv.go
```

## 4. HTTP Server

You can launch an HTTP server by run the following command, that will allow you to receive JSON files with user objects slice through an HTTP request.
Run ``make http-server`` or:

```
go run ./cmd/server/server_http.go
```
