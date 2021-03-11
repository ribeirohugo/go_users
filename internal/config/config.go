package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

const configFile = "config.toml"

type Config struct {
	BinFile string `toml:"bin_path"`

	ConDb       string `toml:"connection_db"`
	ConHost     string `toml:"connection_host"`
	ConPassword string `toml:"connection_password"`
	ConPort     int    `toml:"connection_port"`
	ConUser     string `toml:"connection_user"`
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
