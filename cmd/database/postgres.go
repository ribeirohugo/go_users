package main

import (
	"github.com/ribeirohugo/go_users/internal/config"
	"github.com/ribeirohugo/go_users/internal/fault"
	"github.com/ribeirohugo/go_users/internal/model"
)

func main() {
	cfg, err := config.Load()
	fault.HandleFatalError("", err)

	con := model.PostgresCon{
		Config: model.Config{
			Host:     cfg.PostgresHost,
			Port:     cfg.PostgresPort,
			User:     cfg.PostgresUser,
			Password: cfg.PostgresPassword,
			Db:       cfg.PostgresDb,
		},
	}

	db := con.New()

	_, err = db.Query("SELECT ( * ) FROM land_registry_price_paid_uk")
	fault.HandleFatalError("", err)

	err = db.Ping()
	fault.HandleFatalError("", err)
}
