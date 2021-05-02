package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

const configFile = "config.toml"

type Db struct {
	Db       string `toml:"db"`
	Host     string `toml:"host"`
	Password string `toml:"password"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
}

type Config struct {
	BinPath string `toml:"bin_path"`
	CsvPath string `toml:"csv_path"`

	HttpHost string `toml:"http_host"`
	TcpHost  string `toml:"tcp_host"`

	MySql    Db `toml:"mysql"`
	Postgres Db `toml:"postgres"`
}

func Load(filePath string) (Config, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	_ = file.Close()

	config := Config{
		BinPath:  "users.bin",
		CsvPath:  "users.csv",
		MySql:    Db{},
		Postgres: Db{},
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
