package config

import (
	"github.com/spf13/viper"
	"rxdrag.com/entify-schema-registry/consts"
)

var c config

type config struct {
	v *viper.Viper
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func Init() {
	c.v = viper.New()
	c.v.SetEnvPrefix(consts.CONFIG_PREFIX)
	c.v.BindEnv(consts.DB_DRIVER)
	c.v.BindEnv(consts.DB_USER)
	c.v.BindEnv(consts.DB_PASSWORD)
	c.v.BindEnv(consts.DB_HOST)
	c.v.BindEnv(consts.DB_PORT)
	c.v.BindEnv(consts.DB_SCHEMA)
	c.v.BindEnv(consts.INSTALLED)
}

func GetString(key string) string {
	return c.v.Get(key).(string)
}

func GetBool(key string) bool {
	return c.v.Get(key) == TRUE
}

func SetString(key string, value string) {
	c.v.BindEnv(key, value)
}

func SetBool(key string, value bool) {
	if value {
		c.v.BindEnv(key, TRUE)
	} else {
		c.v.BindEnv(key, FALSE)
	}
}
