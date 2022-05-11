package repository

import (
	"fmt"

	"rxdrag.com/entify-schema-registry/config"
)

func DbString(cfg config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetString(cfg.User),
		config.GetString(cfg.Password),
		config.GetString(cfg.Host),
		config.GetString(cfg.Port),
		config.GetString(cfg.Database),
	)
}
