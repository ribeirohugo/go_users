package main

import (
	"github.com/ribeirohugo/go_users/internal/config"

	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func main() {
	cfg, err := config.Load()
	fault.HandleFatalError("", err)

	con := model.MysqlCon{
		Config: model.Config{
			Host:     cfg.MysqlDb,
			Port:     cfg.MysqlPort,
			User:     cfg.MysqlUser,
			Password: cfg.MysqlPassword,
			Db:       cfg.MysqlDb,
		},
	}

	db := con.New()

	err = db.Ping()
	fault.HandleFatalError("", err)

	db.Close()
}
