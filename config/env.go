package config

import (
	"os"

	"github.com/spf13/viper"
	"rxdrag.com/entify-schema-registry/consts"
)

var c config

type config struct {
	v *viper.Viper
}

type DbConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func Init() {
	c.v = viper.New()
	c.v.SetEnvPrefix(consts.DB_CONFIG_PREFIX)
	c.v.BindEnv(consts.DB_DRIVER)
	c.v.BindEnv(consts.DB_USER)
	c.v.BindEnv(consts.DB_PASSWORD)
	c.v.BindEnv(consts.DB_HOST)
	c.v.BindEnv(consts.DB_PORT)
	c.v.BindEnv(consts.DB_DATABASE)
	c.v.BindEnv(consts.INSTALLED)
}

func GetString(key string) string {
	return c.v.Get(key).(string)
}

func GetBool(key string) bool {
	return c.v.Get(key) == TRUE
}

func SetString(key string, value string) {
	os.Setenv(consts.DB_CONFIG_PREFIX+"_"+key, value)
}

func SetBool(key string, value bool) {
	if value {
		SetString(key, TRUE)
	} else {
		SetString(key, FALSE)
	}
}

func SetDbConfig(cfg DbConfig) {
	SetString(consts.DB_DRIVER, cfg.Driver)
	SetString(consts.DB_DATABASE, cfg.Database)
	SetString(consts.DB_HOST, cfg.Host)
	SetString(consts.DB_PORT, cfg.Port)
	SetString(consts.DB_USER, cfg.User)
	SetString(consts.DB_PASSWORD, cfg.Password)
}

func GetDbConfig() DbConfig {
	var cfg DbConfig
	cfg.Driver = GetString(consts.DB_DRIVER)
	cfg.Database = GetString(consts.DB_DATABASE)
	cfg.Host = GetString(consts.DB_HOST)
	cfg.Port = GetString(consts.DB_PORT)
	cfg.User = GetString(consts.DB_USER)
	cfg.Password = GetString(consts.DB_PASSWORD)
	return cfg
}
