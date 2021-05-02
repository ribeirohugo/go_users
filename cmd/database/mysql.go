package main

import (
	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

const configFile = "config.toml"

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleFatalError("", err)

	con := model.MysqlCon{
		Config: model.Config{
			Host:     cfg.MySql.Host,
			Port:     cfg.MySql.Port,
			User:     cfg.MySql.User,
			Password: cfg.MySql.Password,
			Db:       cfg.MySql.Db,
		},
	}

	db := con.New()

	err = db.Ping()
	fault.HandleFatalError("", err)

	db.Close()
}
