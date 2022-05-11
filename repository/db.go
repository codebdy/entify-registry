package repository

import (
	"fmt"

	"rxdrag.com/entify-schema-registry/config"
	"rxdrag.com/entify-schema-registry/consts"
)

func DbConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetString(consts.DB_USER),
		config.GetString(consts.DB_PASSWORD),
		config.GetString(consts.DB_HOST),
		config.GetString(consts.DB_PORT),
		config.GetString(consts.DB_DATABASE),
	)
}
