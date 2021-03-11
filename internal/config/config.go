package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

const configFile = "config.toml"

type Config struct {
	BinFile string `toml:"bin_path"`

	HttpHost string `toml:"http_host"`
	HttpPort int    `toml:"http_port"`

	MysqlDb       string `toml:"mysql_db"`
	MysqlHost     string `toml:"mysql_host"`
	MysqlPassword string `toml:"mysql_password"`
	MysqlPort     int    `toml:"mysql_port"`
	MysqlUser     string `toml:"mysql_user"`

	PostgresDb       string `toml:"postgres_db"`
	PostgresHost     string `toml:"postgres_host"`
	PostgresPassword string `toml:"postgres_password"`
	PostgresPort     int    `toml:"postgres_port"`
	PostgresUser     string `toml:"postgres_user"`

	TcpHost string `toml:"tcp_host"`
	TcpPort int    `toml:"tcp_port"`
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

	var config Config
	err = toml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
