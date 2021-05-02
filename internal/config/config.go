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

type Host struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Config struct {
	BinPath string `toml:"bin_path"`
	CsvPath string `toml:"csv_path"`

	MySql    Db `toml:"mysql"`
	Postgres Db `toml:"postgres"`

	HttpHost Host `toml:"http_host"`
	TcpHost  Host `toml:"tcp_host"`
}

func Load() (Config, error) {
	filePath := configFile

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
		MySql:    Db{},
		Postgres: Db{},
		HttpHost: Host{},
		TcpHost:  Host{},
	}

	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
