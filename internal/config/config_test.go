package config

import (
	"io/ioutil"
	"os"
	"testing"
)

var configContent = `bin_path = "users.bin"
csv_path = "users.csv"
http_host = "localhost:8080"
tcp_host = "localhost:8081"

[mysql]
db = "mysql_db"
host = "mysql_host"
password = "mysql_password"
port = 3306
user = "mysql_user"

[postgres]
db = "postgres_db"
host = "postgres_host"
password = "postgres_password"
port = 3306
user = "postgres_user"
`

var configTest = Config{
	BinPath:  "users.bin",
	CsvPath:  "users.csv",
	HttpHost: "localhost:8080",
	TcpHost:  "localhost:8081",
	MySql: Db{
		Db:       "mysql_db",
		Host:     "mysql_host",
		Password: "mysql_password",
		Port:     3306,
		User:     "mysql_user",
	},
	Postgres: Db{
		Db:       "postgres_db",
		Host:     "postgres_host",
		Password: "postgres_password",
		Port:     3306,
		User:     "postgres_user",
	},
}

func TestConfig(t *testing.T) {

	tempFile, _ := createTempFile()

	defer os.Remove(tempFile.Name())

	cfg, _ := Load(tempFile.Name())

	if cfg != configTest {
		t.Errorf("Wrong config file output,\n got: %v,\n want: %v.", cfg, configTest)
	}

	tempFile.Close()
}

func createTempFile() (*os.File, error) {

	tempFile, err := ioutil.TempFile("", configFile)
	if err != nil {
		return nil, err
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(configContent)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
