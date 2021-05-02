package main

import (
	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func main() {
	cfg, err := config.Load(configFile)
	fault.HandleFatalError("", err)

	con := model.PostgresCon{
		Config: model.Config{
			Host:     cfg.Postgres.Host,
			Port:     cfg.Postgres.Port,
			User:     cfg.Postgres.User,
			Password: cfg.Postgres.Password,
			Db:       cfg.Postgres.Db,
		},
	}

	db := con.New()

	_, err = db.Query("SELECT ( * ) FROM column")
	fault.HandleFatalError("", err)

	err = db.Ping()
	fault.HandleFatalError("", err)
}
